package gkin

import (
	"github.com/kinyz/gkin/bucket"
	"github.com/kinyz/gkin/consumer"
	"github.com/kinyz/gkin/pb"
	"github.com/kinyz/gkin/producer"
	"github.com/kinyz/gkin/storage"
)

type (
	Message = *pb.Message

	Producer = producer.Producer

	Consumer = consumer.Consumer

	Bucket = *bucket.Bucket

	Storage = storage.Storage
)

func NewBucket(key string, sto Storage) Bucket {
	return bucket.NewBucket(key, sto)
}

func NewProducer(addr, clientId, key string) (Producer, error) {

	return producer.NewProducer(addr, clientId, key)
}

func NewConsumer(addr, key string) Consumer {
	return consumer.NewConsumer(addr, key)
}

func LocalStorage() Storage {
	return storage.NewLocalStorage()
}
