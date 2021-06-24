package bucket

import (
	"context"
	"errors"
	"fmt"
	"gkin/channel"
	"gkin/pb"
	"gkin/pool"
	"gkin/storage"
	"gkin/utils"
	"google.golang.org/grpc"
	"log"
	"net"
	"sync"
	"time"
)

func NewBucket(sto storage.Storage) *Bucket {
	return &Bucket{
		sto:          sto,
		consumerChan: channel.NewChannel(100000),
		topicMgr:     newTopicManager(),
		connMgr:      newConnectManager()}
}

type Bucket struct {
	sto     storage.Storage
	addr    string
	workId  int
	srv     *grpc.Server
	connMgr *connectManager

	consumerChan *channel.Channel

	topicMgr *topicManage

	maxConnNum int
	nowConnNum int
	connLock   sync.RWMutex
}

const StreamKey = "dwigoqhdoiq(U)(J()_"

func (b *Bucket) RequestConnect(ctx context.Context, conn *pb.Connection) (*pb.Connection, error) {
	if conn.GetKey() != StreamKey {
		return nil, errors.New("验证key错误")
	}
	conn.Token = utils.NewToken()
	b.connMgr.AddOrUpData(conn)
	return conn, nil
}

func (b *Bucket) SetMaxConnNum(max int) {
	b.maxConnNum = max
}

func (b *Bucket) WatchStream(req *pb.RequestListenTopic, server pb.Stream_WatchStreamServer) error {
	if !b.connMgr.Oauth(req.GetConn().GetClientId(), req.GetConn().GetToken()) {
		return errors.New("token验证失败")
	}
	b.addConnNum()
	b.connMgr.AddWatcher(req.GetConn().GetClientId(), server)
	log.Println(req.GetConn().GetClientId(), " 加入监听 ")
	for k, v := range req.GetTopicGroup() {
		t, err := b.topicMgr.Get(k)
		if err != nil {
			panic(err)
		}
		t.AddGroup(v, req.GetConn().GetClientId())
		log.Println("监听topic :", k, " 组别:", v)
	}
	log.Println("当前连接数:", b.getConnNum())
	<-server.Context().Done()

	b.reduceConnNum()
	b.connMgr.RemoveWatcher(req.GetConn().GetClientId())

	return nil
}

func (b *Bucket) SendStream(conn pb.Stream_SendStreamServer) error {
	if b.getConnMaxNum() > 0 && b.getConnNum() >= b.getConnMaxNum() {
		return errors.New("当前连接数已满")
	}
	ctx, cancel := context.WithCancel(conn.Context())
	b.addConnNum()
	go func() {
		for {
			recv, err := conn.Recv()
			if err != nil {
				cancel()
				break
			}
			if !b.connMgr.Oauth(recv.GetConn().GetClientId(), recv.GetConn().GetToken()) {
				cancel()
				break
			}
			message, err := b.saveMessage(recv.GetMessage())
			if err != nil {
				go func() {
					log.Println("存储错误")
					res := pool.GetMessageResponse()
					res.Uuid = recv.GetMessage().GetUuid()
					res.Result = false
					conn.Send(res)
					pool.PutMessageResponse(res)
				}()
				continue
			} else {
				go func() {
					res := pool.GetMessageResponse()
					res.Uuid = recv.GetMessage().GetUuid()
					res.Result = true
					res.Sequence = message.GetSequence()
					conn.Send(res)
					pool.PutMessageResponse(res)
				}()
			}

			ch, err := b.consumerChan.Get()
			if err != nil {
				log.Println(err)
				continue
			}
			ch <- message

			log.Println("收到消息", message)
		}
	}()

	log.Println("加入消费者  ")
	log.Println("当前连接数:", b.getConnNum())
	<-ctx.Done()
	b.reduceConnNum()
	return nil
}
func (b *Bucket) Stop() {
	b.srv.Stop()
}

func (b *Bucket) watchChannel(workId int) {
	c, err := b.consumerChan.Get()
	if err != nil {
		log.Println("watch ", workId, "启动失败：", err)
		return
	}
	log.Println("watch work ", workId, " 启动成功")
	for {
		select {
		case msg := <-c:
			b.preHandleWatchMessage(msg.(*pb.Message))
		}
	}
}

func (b *Bucket) preHandleWatchMessage(msg *pb.Message) {
	t, err := b.topicMgr.Get(msg.GetTopic())
	if err != nil {
		return
	}
	i := 0
	for _, v := range t.GetGroups() {
		err := b.sendWatch(v, msg)
		if err != nil {
			continue
		}
		i++
	}

	if i > 1 {
		msg.IsConsume = true
		_, err := b.saveMessage(msg)
		if err != nil {
			log.Println("存盘错误", err)
		}
	}
}
func (b *Bucket) sendWatch(list *pb.TopicGroup, msg *pb.Message) error {

	for k, _ := range list.Followers {
		err := b.connMgr.SendWatcher(k, msg)
		if err != nil {
			continue
		}
		return nil
	}
	return errors.New("不存在消费者")

}

func (b *Bucket) start() error {

	b.sto.Initialization()

	ln, err := net.Listen("tcp", b.addr)
	if err != nil {
		fmt.Println("网络异常", err)
		return err
	}
	b.srv = grpc.NewServer()
	pb.RegisterStreamServer(b.srv, b)
	log.Println("bucket启动成功.......")
	log.Println(b.addr)

	for i := 0; i < 10; i++ {
		go b.watchChannel(i)
	}
	err = b.srv.Serve(ln)
	if err != nil {
		return err
	}
	return nil
}
func (b *Bucket) Serve(addr string) {
	b.addr = addr
	err := b.start()
	if err != nil {
		panic(err)
	}
}

func (b *Bucket) saveMessage(msg *pb.Message) (*pb.Message, error) {
	msg.IsWrite = true
	t, _ := b.topicMgr.AddSequence(msg.GetTopic())
	msg.Sequence = t.GetLastSequence()
	msg.TimesTamp = time.Now().UnixNano()
	err := b.sto.Write(msg)
	if err != nil {
		log.Println(msg.GetSequence(), "写入失败", err)
		return nil, err
	}
	return msg, nil
}
