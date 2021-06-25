package consumer

import (
	"context"
	"errors"
	"github.com/kinyz/gkin/pb"
	"github.com/kinyz/gkin/utils"
	"google.golang.org/grpc"
	"log"
)

type Consumer interface {
	Watches(topic []string, group string, f func(message *pb.Message)) error
	//Reconnection() error
	Watch(topic string, group string, f func(message *pb.Message)) error
}

func NewConsumer(addr, key string) Consumer {
	//c :=
	//
	//err := c.connect()
	//if err != nil {
	//	return nil, err
	//}
	return &consumer{addr: addr, key: key}
}

type consumer struct {
	addr string

	//head *pb.Connection

	key string
	//cli pb.StreamClient
}

func (c *consumer) connect() (pb.StreamClient, *pb.Connection, error) {
	conn, err := c.getConn()
	if err != nil {
		return nil, nil, err
	}
	cli := pb.NewStreamClient(conn)
	connection, err := cli.RequestConnect(context.Background(), &pb.Connection{
		ClientId: utils.NewUuid(),

		Token: "",
		Key:   c.key,
	})

	return cli, connection, nil
}

// Watches 批量监听
func (c *consumer) Watches(topics []string, group string, f func(message *pb.Message)) error {

	if len(topics) < 1 {
		return errors.New("topic不能为空")
	}

	cli, head, err := c.connect()
	if err != nil {
		return err
	}
	ch := make(chan *pb.Message)
	t := make(map[string]string)
	for _, v := range topics {
		t[v] = group
	}
	ctx, cancel := context.WithCancel(context.Background())
	go c.preHandleRecv(ch, f, ctx)
	log.Println("topics:", topics, " group:", group, " watching....")
	err = c.watch(t, ch, cli, head)
	if err != nil {
		log.Println(err)
	}
	close(ch)
	cancel()

	return err
}

// Watch 批量监听
func (c *consumer) Watch(topic string, group string, f func(message *pb.Message)) error {
	if len(topic) < 1 {
		return errors.New("topic不能为空")
	}

	cli, head, err := c.connect()
	if err != nil {
		return err
	}
	ch := make(chan *pb.Message)
	t := make(map[string]string)
	t[topic] = group
	ctx, cancel := context.WithCancel(context.Background())
	go c.preHandleRecv(ch, f, ctx)
	log.Println("topic:", topic, " group:", group, " watching....")
	err = c.watch(t, ch, cli, head)
	if err != nil {
		log.Println(err)
	}
	close(ch)
	cancel()

	return err

}

func (c *consumer) preHandleRecv(ch chan *pb.Message, f func(message *pb.Message), ctx context.Context) {
	for {
		select {
		case msg := <-ch:
			f(msg)
		case <-ctx.Done():
			//log.Println("关闭通道")
			return
		}
	}

}

func (c *consumer) watch(topic map[string]string, ch chan *pb.Message, cli pb.StreamClient, head *pb.Connection) error {

	stream, err := cli.WatchStream(context.Background(), &pb.RequestListenTopic{
		Conn:       head,
		TopicGroup: topic,
	})
	if err != nil {
		return err
	}
	//	stream.
	//ch:=make(chan *pb.Message)
	for {
		msg, err := stream.Recv()
		if err != nil {
			break
		}
		ch <- msg
	}

	return nil
}

//func (c *consumer) Reconnection() error {
//	return c.connect()
//}
func (c *consumer) getConn() (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(c.addr, grpc.WithInsecure())
	if err != nil {
		log.Println("连接服务器失败", err)
		return nil, err
	}
	return conn, nil
}
