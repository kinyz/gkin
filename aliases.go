package gkin

import (
	"gkin/pb"
	"gkin/producer"
)

type (
	Message = *pb.Message

	Producer = producer.Producer
)

func NewProducer(addr, clientId, key string) (Producer, error) {

	return producer.NewProducer(addr, clientId, key)
}
