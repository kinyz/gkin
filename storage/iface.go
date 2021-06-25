package storage

import (
	"github.com/kinyz/gkin/pb"
)

type Storage interface {
	Initialization()
	Get(topic string, sequence int64) (*pb.Message, error)
	Write(msg *pb.Message) error
	IsExist(topic string, sequence int64) bool
	Len(topic string) (int, error)
	//GetSequences(n)
}
