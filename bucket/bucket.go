package bucket

import (
	"context"
	"errors"
	"fmt"
	"github.com/kinyz/gkin/pb"
	"github.com/kinyz/gkin/pool"
	"github.com/kinyz/gkin/storage"
	"github.com/kinyz/gkin/utils"
	"google.golang.org/grpc"
	"log"
	"net"
	"sync"
	"sync/atomic"
	"time"
)

func NewBucket(key string, sto storage.Storage) *Bucket {
	return &Bucket{
		sto:             sto,
		consumerChan:    make(chan *pb.Message),
		topicMgr:        newTopicManager(),
		connMgr:         newConnectManager(),
		storageChan:     make(chan *pb.Message),
		topicMessageNum: make(map[string]int64),
		key:             key,
	}
}

type Bucket struct {
	sto     storage.Storage
	addr    string
	workId  int
	srv     *grpc.Server
	connMgr *connectManager

	//consumerChan *channel.Channel
	consumerChan chan *pb.Message
	topicMgr     *topicManage

	maxConnNum int64
	nowConnNum int64
	connLock   sync.RWMutex

	storageChan chan *pb.Message

	topicMessageNum map[string]int64

	key string
}

//const StreamKey = "dwigoqhdoiq(U)(J()_"

func (b *Bucket) RequestConnect(ctx context.Context, conn *pb.Connection) (*pb.Connection, error) {
	if conn.GetKey() != b.key {
		return nil, errors.New("验证key错误")
	}
	conn.Token = utils.NewToken()
	b.connMgr.AddOrUpData(conn)
	return conn, nil
}

func (b *Bucket) SetMaxConnNum(max int64) {
	//b.maxConnNum = max
	atomic.SwapInt64(&b.maxConnNum, max)
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
	log.Println(req.GetConn().GetClientId(), " 离开监听")
	log.Println("当前连接数:", b.getConnNum())
	return nil
}

func (b *Bucket) SendStream(conn pb.Stream_SendStreamServer) error {
	if b.getConnMaxNum() > 0 && b.getConnNum() >= b.getConnMaxNum() {
		return errors.New("当前连接数已满")
	}

	c := pool.GetConnect()
	err := conn.RecvMsg(c)
	if err != nil {
		return err
	}
	if !b.connMgr.Oauth(c.GetClientId(), c.GetToken()) {
		return errors.New("token验证失败")
	}

	ctx, cancel := context.WithCancel(conn.Context())
	b.addConnNum()
	log.Println("加入生产者  ")
	log.Println("当前连接数:", b.getConnNum())
	//go func() {
	for {
		recv, err := conn.Recv()
		if err != nil {
			cancel()
			break
		}

		//b.topicMessageNum[]

		t, _ := b.topicMgr.AddSequence(recv.GetTopic())
		//i:=b.topicMessageNum[recv.GetMessage().GetTopic()]
		//atomic.AddInt64(&i,1)

		recv.Sequence = t.GetLastSequence()
		recv.TimesTamp = time.Now().UnixNano()
		recv.Producer = c.GetClientId()

		log.Println("收到消息", recv)
		//if !b.connMgr.Oauth(recv.GetConn().GetClientId(), recv.GetConn().GetToken()) {
		//	cancel()
		//	break
		//}
		b.storageChan <- recv

		if recv.GetResult() {
			conn.Send(&pb.ResponseMessage{
				Uuid:     recv.GetUuid(),
				Sequence: recv.GetSequence(),
				Result:   true,
			})
		}
		//select {
		//case b.consumerChan<-recv:
		//	//log.Println("收到消息", recv)
		//case b.storageChan<-recv:
		//
		//}

	}
	//	}()

	<-ctx.Done()
	b.reduceConnNum()
	log.Println("生产者离开 ")
	log.Println("当前连接数:", b.getConnNum())
	err = b.connMgr.Remove(c.GetClientId())
	if err != nil {
		log.Println(err)
	}
	return nil
}
func (b *Bucket) Stop() {
	b.srv.Stop()
}

func (b *Bucket) watchChannel(workId int) {
	//c, err := b.consumerChan.Get()
	//if err != nil {
	//	log.Println("watch ", workId, "启动失败：", err)
	//	return
	//}
	log.Println("watch work ", workId, " 启动成功")
	for {
		select {
		case msg := <-b.consumerChan:
			log.Println(workId, "发送消费消息", msg.GetSequence())
			b.preHandleWatchMessage(msg)
			//_, err := b.saveMessage(msg)
			//if err != nil {
			//	//go func() {
			//	//	log.Println("存储错误")
			//	//	res := pool.GetMessageResponse()
			//	//	res.Uuid = recv.GetMessage().GetUuid()
			//	//	res.Result = false
			//	//	conn.Send(res)
			//	//	pool.PutMessageResponse(res)
			//	//}()
			//	log.Println("存储错误:",err)
			//	//continue
			//}

		}
	}
}

func (b *Bucket) preHandleWatchMessage(msg *pb.Message) {
	t, err := b.topicMgr.Get(msg.GetTopic())
	if err != nil {
		return
	}
	//	i := 0
	for _, v := range t.GetGroups() {
		err := b.sendWatch(v, msg)
		if err != nil {
			continue
		}
		//i++
	}
	//
	//go func() {
	//	if i > 1 {
	//		msg.IsConsume = true
	//		_, err := b.saveMessage(msg)
	//		if err != nil {
	//			log.Println("存盘错误", err)
	//		}
	//	}
	//}()

}
func (b *Bucket) sendWatch(list *pb.TopicGroup, msg *pb.Message) error {

	for k, _ := range list.Followers {
		err := b.connMgr.SendWatcher(k, msg)
		if err != nil {
			continue
		}

	}
	//log.Println("发送消费消息",msg.GetSequence())

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
		go b.storageChannel(i)
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

func (b *Bucket) storageChannel(workId int) {
	log.Println("storage work ", workId, " 启动成功")
	for {
		select {
		case msg := <-b.storageChan:

			log.Println(workId, "收到存储消息", msg.GetSequence())
			b.consumerChan <- msg
			b.saveMessage(msg)
		}
	}
}
func (b *Bucket) saveMessage(msg *pb.Message) (*pb.Message, error) {
	msg.IsWrite = true
	err := b.sto.Write(msg)
	if err != nil {
		log.Println(msg.GetSequence(), "写入失败", err)
		return nil, err
	}
	return msg, nil
}
