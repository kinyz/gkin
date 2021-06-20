package main

import (
	"context"
	"gkin/message"
	"gkin/utils"
	"google.golang.org/grpc"
	"log"
	"strconv"
)

func main(){

	conn, err := getConn()
	if err != nil {
		return
	}

	c:=message.NewMessageStreamClient(conn)
	connect, err := c.RequestConnect(context.Background(), &message.Connect{
		ClientId: utils.NewUuid(),
		Oauth:    false,
		Token:    "",
		Key:      "",
	})
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("验证成功",connect)

	msg:=message.New()
	msg.SetTopic("user")
	msg.SetKey("im key")
	for i:=0;i<10;i++ {
		msg.SetValue([]byte("我是value"+strconv.Itoa(i)))
		send, err := c.SyncSend(context.Background(), msg)
		if err != nil {
			log.Println("发送失败",err)
			continue
		}
		log.Println(" 发送结果",send)

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