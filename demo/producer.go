package main

import (
	"context"
	"gkin/pb"
	"gkin/pool"
	"time"

	"gkin/utils"
	"google.golang.org/grpc"
	"log"
	"strconv"
)

func main() {

	conn, err := getConn()
	if err != nil {
		return
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

	msg := pool.GetMessage()
	msg.SetTopic("user")
	msg.SetKey("im key")
	cli, err := c.SendStream(context.Background())
	if err != nil {
		panic(err)

	}
	for i := 0; i < 10; i++ {
		msg.SetValue([]byte("我是value" + strconv.Itoa(i)))
		err := cli.Send(&pb.RequestSendStream{
			Conn:    connection,
			Message: msg,
		})
		if err != nil {
			log.Println(err)
			continue
		}

		log.Println(" 发送成功")
		time.Sleep(1 * time.Second)

	}

	conn.Close()

}

func getConn() (*grpc.ClientConn, error) {
	conn, err := grpc.Dial("127.0.0.1:17222", grpc.WithInsecure())
	if err != nil {
		log.Println("连接服务器失败", err)
		return nil, err
	}
	return conn, nil
}
