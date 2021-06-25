package main

import (
	"context"
	"gkin/pb"
	"gkin/utils"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := getWatchConn()
	if err != nil {
		panic(err)
	}
	c := pb.NewStreamClient(conn)
	connection, err := c.RequestConnect(context.Background(), &pb.Connection{
		ClientId: utils.NewUuid(),

		Token: "",
		Key:   "dwigoqhdoiq(U)(J()_",
	})
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("验证成功", connection)

	listen := make(map[string]string)
	listen["user"] = "1112"

	stream, err := c.WatchStream(context.Background(), &pb.RequestListenTopic{
		Conn:       connection,
		TopicGroup: listen,
	})
	if err != nil {
		return
	}

	for {
		recv, err := stream.Recv()
		if err != nil {
			return
		}
		log.Println("监听到消息:", recv.GetSequence())
		log.Println(recv)
	}
}

func getWatchConn() (*grpc.ClientConn, error) {
	//conn, err := grpc.Dial("1.116.248.126:17222", grpc.WithInsecure())
	conn, err := grpc.Dial("127.0.0.1:17222", grpc.WithInsecure())
	if err != nil {
		log.Println("连接服务器失败", err)
		return nil, err
	}
	return conn, nil
}
