package producer

import (
	"context"
	"errors"
	"gkin/pb"
	"gkin/pool"
	"gkin/utils"
	"google.golang.org/grpc"
	"log"
	"time"
)

type Producer interface {
	// SyncSend 同步发送
	SyncSend(topic, key string, value []byte, headers ...map[string]string) (bool, error)
	// ASyncSend 异步发送
	ASyncSend(topic, key string, value []byte, headers ...map[string]string) error
	// Close 关闭连接
	Close() error
	// Reconnection 重新连接
	Reconnection() error
}

func NewProducer(addr, clientId, key string) (Producer, error) {
	p := &producer{syncMap: make(map[string]chan bool),
		addr: addr,
		head: &pb.Connection{
			ClientId: clientId,
			Token:    "",
			Key:      key,
		}}
	err := p.connect()
	if err != nil {
		return nil, err
	}
	go p.responseChannel()
	return p, nil
}

type producer struct {
	head    *pb.Connection
	cli     pb.Stream_SendStreamClient
	syncMap map[string]chan bool
	addr    string
}

func (p *producer) send(uuid, topic, key string, value []byte, headers ...map[string]string) error {
	msg := pool.GetMessage()

	msg.Uuid = uuid
	msg.SetKey(key)
	msg.SetValue(value)
	msg.SetTopic(topic)
	if len(headers) > 0 {
		msg.SetHeaders(headers[0])
	}
	err := p.cli.Send(msg)
	if err != nil {
		pool.PutMessage(msg)
		return err
	}
	pool.PutMessage(msg)
	return nil
}

// SyncSend 同步发送
func (p *producer) SyncSend(topic, key string, value []byte, headers ...map[string]string) (bool, error) {
	uuid := utils.NewUuid()
	p.syncMap[uuid] = make(chan bool)
	if err := p.send(uuid, topic, key, value, headers...); err != nil {
		close(p.syncMap[uuid])
		delete(p.syncMap, uuid)
		return false, err
	}
	select {
	case b := <-p.syncMap[uuid]:
		close(p.syncMap[uuid])
		delete(p.syncMap, uuid)
		return b, nil
	case <-time.After(20 * time.Second):
		break
	}
	close(p.syncMap[uuid])
	delete(p.syncMap, uuid)
	return false, errors.New("timeout")
}

// ASyncSend 异步发送
func (p *producer) ASyncSend(topic, key string, value []byte, headers ...map[string]string) error {
	return p.send(utils.NewUuid(), topic, key, value, headers...)
}
func (p *producer) responseChannel() {
	for {
		recv, err := p.cli.Recv()
		if err != nil {
			return
		}
		syncChan, ok := p.syncMap[recv.GetUuid()]
		if ok {
			syncChan <- recv.Result
		}
	}

}
func (p *producer) connect() error {
	conn, err := p.getConn()
	if err != nil {
		return err
	}
	c := pb.NewStreamClient(conn)
	connection, err := c.RequestConnect(context.Background(), &pb.Connection{
		ClientId: p.head.GetClientId(),

		Token: "",
		Key:   p.head.GetKey(),
	})
	if err != nil {
		log.Println(err)
		return err
	}
	p.head = connection

	stream, err := c.SendStream(context.Background())
	if err != nil {
		return err
	}
	err = stream.SendMsg(p.head)
	if err != nil {
		return err
	}

	p.cli = stream

	return nil
}
func (p *producer) Close() error {
	return p.cli.CloseSend()
}
func (p *producer) Reconnection() error {
	return p.connect()
}
func (p *producer) getConn() (*grpc.ClientConn, error) {
	//conn, err := grpc.Dial("1.116.248.126:17222", grpc.WithInsecure())
	conn, err := grpc.Dial(p.addr, grpc.WithInsecure())

	if err != nil {
		log.Println("连接服务器失败", err)
		return nil, err
	}
	return conn, nil
}
