package bucket

import (
	"errors"
	"gkin/pb"
	"gkin/pool"
	"log"
	"sync"
)

func newTopicManager() *topicManage {
	return &topicManage{TopicList: &pb.TopicList{
		List: make(map[string]*pb.Topic),
	}}
}

type topicManage struct {
	*pb.TopicList
	lock sync.RWMutex
}

func (mgr *topicManage) Add(topic string) (*pb.Topic, error) {
	t, ok := mgr.List[topic]
	if ok {
		return t, errors.New(topic + "已存在")
	}
	to := pool.GetTopic()
	to.Name = topic
	log.Println("新增topic", topic)
	mgr.List[topic] = to
	return mgr.List[topic], nil
}
func (mgr *topicManage) Get(topic string) (*pb.Topic, error) {
	mgr.lock.RLock()
	defer mgr.lock.RUnlock()
	t, ok := mgr.List[topic]
	if !ok {
		return mgr.Add(topic)
	}
	return t, nil
}
func (mgr *topicManage) UpData(topic *pb.Topic) error {
	_, err := mgr.Get(topic.GetName())
	if err != nil {
		return err
	}
	mgr.lock.Lock()
	mgr.List[topic.GetName()] = topic
	mgr.lock.Unlock()
	return nil
}
func (mgr *topicManage) AddSequence(topic string) (*pb.Topic, error) {
	_, err := mgr.Get(topic)
	if err != nil {
		return nil, err
	}
	//log.Println("我是  ", mgr.List[topic])
	mgr.lock.Lock()
	defer mgr.lock.Unlock()
	mgr.List[topic].LastSequence++
	mgr.List[topic].MessageLen++
	return mgr.List[topic], nil
}
func (mgr *topicManage) Remove(topic string) error {
	t, err := mgr.Get(topic)
	if err != nil {
		return err
	}
	mgr.lock.Lock()
	pool.PutTopic(t)
	delete(mgr.List, topic)
	mgr.lock.Unlock()
	return nil
}
