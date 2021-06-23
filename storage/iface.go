package storage

import (
	"gkin/message"
)

type Storage interface {
	Initialization()
	Get(topic string, sequence int64) (*message.Message, error)
	Write(msg *message.Message) error
	IsExist(topic string, sequence int64) bool

	Len(topic string) (int, error)
	//GetSequences(n)
}
