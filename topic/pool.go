package topic

import (
	"gkin/pb"
	"sync"
)

var Pool = sync.Pool{New: func() interface{} {

	return &pb.Topic{
		Name:         "",
		Groups:       make(map[string]*pb.TopicGroup),
		MessageLen:   0,
		LastSequence: 0,
	}
}}

func New() *pb.Topic {
	m := Pool.Get().(*pb.Topic)
	m.Groups = make(map[string]*pb.TopicGroup)
	m.LastSequence = 0
	m.MessageLen = 0
	return m
}
