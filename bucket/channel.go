package bucket

import (
	"errors"
	"gkin/channel"
	"gkin/pb"
	"sync"
)

type consumerChannel struct {
	topicChan map[string]*channel.Channel
	lock      sync.RWMutex
}

func newConsumerChannel() *consumerChannel {
	return &consumerChannel{topicChan: make(map[string]*channel.Channel)}
}
func (cc *consumerChannel) getChan(topic string) (*channel.Channel, error) {
	c, ok := cc.topicChan[topic]
	if ok {
		return c, nil
	}
	return cc.makeTopicChan(topic)
}

func (cc *consumerChannel) makeTopicChan(topic string) (*channel.Channel, error) {
	cc.lock.Lock()
	defer cc.lock.Unlock()
	_, ok := cc.topicChan[topic]
	if ok {
		return nil, errors.New("通道已存在")
	}
	cc.topicChan[topic] = channel.NewChannel()
	return cc.topicChan[topic], nil
}
func (cc *consumerChannel) remove(topic string) {
	c, err := cc.getChan(topic)
	if err != nil {
		return
	}
	cc.lock.Lock()
	c.Close()
	delete(cc.topicChan, topic)
	cc.lock.Unlock()
}

func (cc *consumerChannel) send(topic string, msg *pb.Message) error {
	c, err := cc.getChan(topic)
	if err != nil {
		return err
	}
	return c.Send(msg)
}

type producerChannel struct {
}
