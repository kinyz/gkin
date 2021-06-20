package bucket

import (
	"context"
	"fmt"
	"gkin/message"
	"gkin/storage"
	"gkin/utils"
	"google.golang.org/grpc"
	"log"
	"net"
	"sync"
	"time"
)

func NewBucket(sto storage.Storage)*Bucket{
	return &Bucket{
	sto: sto,
	producers: make(chan *message.Message),
	consumers: make(map[string]chan *message.Message)}}

type Bucket struct {
	sto storage.Storage
	producers chan *message.Message
	addr string
	workId int

	consumers map[string]chan *message.Message

	sequence int64

	sequenceLock sync.Mutex

}

func (b *Bucket) RequestConnect(ctx context.Context, conn *message.Connect) (*message.Connect, error) {
	//b.producers [conn.GetClientId()] = make(chan *message.Message)
	conn.Oauth = true
	conn.Token = utils.NewToken()

	return conn,nil
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
		b.producers<-m
		return &message.RespSend{
			Sequence:  b.sequence,
			IsConsume: true,
		},nil
}

func (b *Bucket) producerChannel(workId int){

	log.Println("工作通道已开启,workId=",workId)
	for  {
		select {
		case msg:=<-b.producers:
			log.Println("通道",workId,"-------- ")
			log.Println(msg.GetTopic())
			log.Println(msg.GetKey())
			log.Println(msg.GetSequence())
			log.Println(msg.GetTimesTamp())
		}
	}
}

func (b*Bucket)start()error{
	ln, err := net.Listen("tcp", b.addr)
	if err != nil {
		fmt.Println("网络异常", err)
		return err
	}
	s := grpc.NewServer()
	message.RegisterMessageStreamServer(s,b)
	go b.producerChannel(b.workId)
	log.Println("Listen ", b.addr)

	err = s.Serve(ln)
	if err != nil {
		fmt.Println("网络启动异常", err)
		return err
	}

	return nil
}
func (b*Bucket)Serve(addr string){
	b.addr = addr
	err := b.start()
	if err != nil {
		panic(err)
	}
}
