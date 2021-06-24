package bucket

import (
	"context"
	"errors"
	"fmt"
	"gkin/connect"
	"gkin/pb"
	"gkin/storage"
	"gkin/topic"
	"gkin/utils"
	"google.golang.org/grpc"
	"log"
	"net"
	"sync"
)

func NewBucket(sto storage.Storage) *Bucket {
	return &Bucket{
		sto: sto,
		//	producers:    make(chan *pb.Message),
		consumerChan: newConsumerChannel(),
		topicMgr:     topic.NewManager(),
		connMgr:      connect.NewManager()}

}

type Bucket struct {
	sto storage.Storage
	//producers chan *pb.Message
	addr   string
	workId int

	srv     *grpc.Server
	connMgr connect.Manager

	//consumers    map[string]channel *message.Message
	sequence     int64
	sequenceLock sync.Mutex
	consumerChan *consumerChannel

	topicMgr *topic.Manage
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

func (b *Bucket) WatchStream(req *pb.RequestListenTopic, server pb.Stream_WatchStreamServer) error {
	if !b.connMgr.Oauth(req.GetConn().GetClientId(), req.GetConn().GetToken()) {
		return errors.New("token验证失败")
	}
	wg := sync.WaitGroup{}
	wg.Add(1)

	server.Context().Done()
	log.Println(" 我到这了")
	wg.Wait()
	return nil
}

func (b *Bucket) SendStream(conn pb.Stream_SendStreamServer) error {

	go func() {
		for {
			recv, err := conn.Recv()
			if err != nil {
				//	wg.Done()
				break
			}
			if !b.connMgr.Oauth(recv.GetConn().GetClientId(), recv.GetConn().GetToken()) {
				log.Println("收到消息验证错误")
				//	wg.Done()
				break
			}
			log.Println("收到消息", recv.GetMessage())
		}
	}()

	<-conn.Context().Done()
	log.Println("到底2")
	return nil
}
func (b *Bucket) Stop() {
	b.srv.Stop()
}

//func (b *Bucket) SyncSend(ctx context.Context, m *pb.Message) (*pb.RespSend, error) {
//
//	t, err := b.topicMgr.AddSequence(m.Topic)
//	if err != nil {
//		return nil, err
//	}
//	m.Sequence = t.GetLastSequence()
//	m.TimesTamp = time.Now().UnixNano()
//	b.producers <- m
//	return &pb.RespSend{
//		Sequence:  b.sequence,
//		IsConsume: true,
//	}, nil
//}
//
//func (b *Bucket) producerChannel(workId int) {
//
//	log.Println("工作通道已开启,workId=", workId)
//	for {
//		select {
//		case msg := <-b.producers:
//			go b.saveMessage(msg)
//			c, err := b.consumerChan.getChan(msg.GetTopic())
//			if err != nil {
//				log.Println(err)
//				continue
//			}
//			err = c.Send(msg)
//			if err != nil {
//				log.Println("发送通道失败", err)
//			}
//		}
//	}
//}

func (b *Bucket) preHandleConsumerChan(topic string) error {
	c, err := b.consumerChan.getChan(topic)
	if err != nil {
		return err
	}
	ch, err := c.Get()
	if err != nil {
		return err
	}
	log.Println("开始监听")
	for {
		select {
		case msg := <-ch:
			log.Println("收到消费者信息", msg.(*pb.Message))
		case <-c.Done():
			log.Println("通道已关闭")
			break

		}

	}
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
	//go b.producerChannel(b.workId)

	//	go b.preHandleConsumerChan("user")

	log.Println("stream启动成功.......")
	log.Println(b.addr)
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

func (b *Bucket) saveMessage(msg *pb.Message) error {
	err := b.sto.Write(msg)
	if err != nil {
		log.Println(msg.GetSequence(), "写入失败", err)
		return err
	}
	return nil

}
