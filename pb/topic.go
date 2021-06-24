package pb

import (
	"errors"
	"sync"
)

var groupLock sync.RWMutex

// AddGroup 增加监听组成员
func (x *Topic) AddGroup(name, uuid string) {
	groupLock.Lock()

	_, ok := x.Groups[name]
	if !ok {
		x.Groups[name] = &TopicGroup{Followers: make(map[string]string)}

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

func (x *Topic) AddMessage(id int64) {
	x.Messages[id] = false
}
func (x *Topic) GetMessageState(id int64) (bool, error) {
	b, ok := x.Messages[id]
	if !ok {
		return false, errors.New("msg不存在")
	}
	return b, nil
}

func (x *Topic) SetMessageState(id int64, b bool) error {
	b, ok := x.Messages[id]
	if !ok {
		return errors.New("msg不存在")
	}
	x.Messages[id] = b
	return nil

}
