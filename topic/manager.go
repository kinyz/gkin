package topic

import (
	"errors"
	"log"
	"sync"
)

func NewManager() *Manage {
	return &Manage{TopicList: &TopicList{
		List: make(map[string]*Topic),
	}}
}

type Manage struct {
	*TopicList
	lock sync.RWMutex
}

func (mgr *Manage) Add(topic string) (*Topic, error) {
	t, ok := mgr.List[topic]
	if ok {
		return t, errors.New(topic + "已存在")
	}
	to := New()
	to.Name = topic
	log.Println("新增topic", topic)
	mgr.List[topic] = to
	return mgr.List[topic], nil
}
func (mgr *Manage) Get(topic string) (*Topic, error) {
	mgr.lock.RLock()
	defer mgr.lock.RUnlock()
	t, ok := mgr.List[topic]
	if !ok {
		return mgr.Add(topic)
	}
	return t, nil
}
func (mgr *Manage) UpData(topic *Topic) error {
	_, err := mgr.Get(topic.GetName())
	if err != nil {
		return err
	}
	mgr.lock.Lock()
	mgr.List[topic.GetName()] = topic
	mgr.lock.Unlock()
	return nil
}
func (mgr *Manage) AddSequence(topic string) (*Topic, error) {
	_, err := mgr.Get(topic)
	if err != nil {
		return nil, err
	}
	log.Println("我是  ", mgr.List[topic])
	mgr.lock.Lock()
	defer mgr.lock.Unlock()
	mgr.List[topic].LastSequence++
	mgr.List[topic].MessageLen++

	return mgr.List[topic], nil
}
func (mgr *Manage) Remove(topic string) error {
	t, err := mgr.Get(topic)
	if err != nil {
		return err
	}
	mgr.lock.Lock()
	Pool.Put(t)
	delete(mgr.List, topic)
	mgr.lock.Unlock()
	return nil
}

var groupLock sync.RWMutex

// AddGroup 增加监听组成员
func (x *Topic) AddGroup(name, uuid string) {
	groupLock.Lock()
	_, ok := x.Groups[name]
	if !ok {
		x.Groups[name].Followers = make(map[string]string)
	}
	x.Groups[name].Followers[uuid] = uuid
	groupLock.Unlock()
}

// GetFollowers 获取监听者列表
func (x *Topic) GetFollowers(name string) (map[string]string, error) {
	groupLock.RLock()
	defer groupLock.RUnlock()
	_, ok := x.Groups[name]
	if !ok {
		return nil, errors.New("Group " + name + " 不存在")
	}
	return x.Groups[name].Followers, nil
}

// RemoveFollower 删除监听者
func (x *Topic) RemoveFollower(name, uuid string) error {
	groupLock.Lock()
	defer groupLock.Unlock()
	_, ok := x.Groups[name]
	if !ok {
		return errors.New("Group " + name + " 不存在")
	}
	_, ok = x.Groups[name].Followers[uuid]
	if !ok {
		return errors.New("Follower " + uuid + " 不存在")
	}
	delete(x.Groups[name].Followers, uuid)
	return nil
}
