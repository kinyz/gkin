package pool

import (
	"github.com/kinyz/gkin/pb"
	"sync"
)

var topicPool = sync.Pool{New: func() interface{} {

	return &pb.Topic{
		Name:         "",
		Groups:       make(map[string]*pb.TopicGroup),
		MessageLen:   0,
		LastSequence: 0,
		Messages:     make(map[int64]bool),
	}
}}

func GetTopic() *pb.Topic {
	m := topicPool.Get().(*pb.Topic)
	m.Groups = make(map[string]*pb.TopicGroup)
	m.Messages = make(map[int64]bool)
	m.LastSequence = 0
	m.MessageLen = 0
	return m
}

func PutTopic(topic *pb.Topic) {
	topicPool.Put(topic)
}
