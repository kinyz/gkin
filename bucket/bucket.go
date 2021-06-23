package bucket

import (
	"context"
	"fmt"
	"gkin/connect"
	"gkin/message"
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
		producers:    make(chan *message.Message),
		consumerChan: newConsumerChannel()}
}

type Bucket struct {
	sto       storage.Storage
	producers chan *message.Message
	addr      string
	workId    int

	//consumers    map[string]channel *message.Message
	sequence     int64
	sequenceLock sync.Mutex
	consumerChan *consumerChannel
}

func (b *Bucket) RequestConnect(ctx context.Context, conn *connect.Connection) (*connect.Connection, error) {
	//b.producers [conn.GetClientId()] = make(channel *message.Message)
	conn.Oauth = true
	conn.Token = utils.NewToken()
	return conn, nil
}

func (b *Bucket) AsyncSend(server message.MessageStream_AsyncSendServer) error {
	panic("implement me")
}

func (b *Bucket) SyncSend(ctx context.Context, m *message.Message) (*message.RespSend, error) {

	b.sequenceLock.Lock()
	b.sequence++
	b.sequenceLock.Unlock()

	m.Sequence = b.sequence
	m.TimesTamp = time.Now().UnixNano()
	//b.consumers[m.GetTopic()]<-m

	b.producers <- m
	return &message.RespSend{
		Sequence:  b.sequence,
		IsConsume: true,
	}, nil
}

func (b *Bucket) producerChannel(workId int) {

	log.Println("工作通道已开启,workId=", workId)
	for {
		select {
		case msg := <-b.producers:
			go b.saveMessage(msg)
			c, err := b.consumerChan.getChan(msg.GetTopic())
			if err != nil {
				log.Println(err)
				continue
			}
			err = c.Send(msg)
			if err != nil {
				log.Println("发送通道失败", err)
			}
		}
	}
}

func (b *Bucket) preHandleConsumerChan(topic string) error {
	c, err := b.consumerChan.getChan(topic)
	if err != nil {
		return err
	}
	ch, err := c.Get()
	if err != nil {
		return err
	}
	for {
		select {
		case msg := <-ch:
			log.Println("收到消费者信息", msg.(*message.Message))
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
	s := grpc.NewServer()
	message.RegisterMessageStreamServer(s, b)
	go b.producerChannel(b.workId)
	log.Println("Listen ", b.addr)
	err = s.Serve(ln)
	if err != nil {
		fmt.Println("网络启动异常", err)
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

func (b *Bucket) saveMessage(msg *message.Message) error {
	err := b.sto.Write(msg)
	if err != nil {
		log.Println(msg.GetSequence(), "写入失败", err)
		return err
	}
	log.Println(msg.GetSequence(), "写入成功")
	return nil

}
