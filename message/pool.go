package message

import (
	"sync"
)

var Pool = sync.Pool{New: func() interface{} {
	return &Message{
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

func New() *Message {
	m := Pool.Get().(*Message)
	m.Headers = make(map[string]string)
	m.IsConsume = false
	m.IsWrite = false
	m.Producer = ""
	return m
}
