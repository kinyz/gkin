package topic

import (
	"sync"
)

var Pool = sync.Pool{New: func() interface{} {

	return &Topic{
		Name:         "",
		Groups:       make(map[string]*TopicGroup),
		MessageLen:   0,
		LastSequence: 0,
	}
}}

func New() *Topic {
	m := Pool.Get().(*Topic)
	m.Groups = make(map[string]*TopicGroup)
	m.LastSequence = 0
	m.MessageLen = 0
	return m
}
