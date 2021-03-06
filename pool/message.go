package pool

import (
	"github.com/kinyz/gkin/pb"
	"sync"
)

var messagePool = sync.Pool{New: func() interface{} {
	return &pb.Message{
		Topic:     "",
		Headers:   make(map[string]string),
		Key:       "",
		Value:     nil,
		TimesTamp: 0,
		IsConsume: false,
		IsWrite:   false,
		Producer:  "",
	}
}}

func GetMessage() *pb.Message {
	m := messagePool.Get().(*pb.Message)
	m.Headers = make(map[string]string)
	m.IsConsume = false
	m.IsWrite = false
	m.Producer = ""
	return m
}

func PutMessage(msg *pb.Message) {
	messagePool.Put(msg)
}

var messageResponsePool = sync.Pool{New: func() interface{} {
	return &pb.ResponseMessage{}
}}

func GetMessageResponse() *pb.ResponseMessage {
	return messageResponsePool.Get().(*pb.ResponseMessage)
}

func PutMessageResponse(msg *pb.ResponseMessage) {
	messageResponsePool.Put(msg)
}
