package bucket

import (
	"errors"
	"github.com/kinyz/gkin/pb"
	"sync"
)

type connectManager struct {
	lock    sync.RWMutex
	connMap sync.Map

	connSend sync.Map

	watcherMap sync.Map
}

func newConnectManager() *connectManager {
	return &connectManager{}
}
func (mgr *connectManager) Get(id string) (*pb.Connection, error) {

	x, ok := mgr.connMap.Load(id)
	if !ok {
		return nil, errors.New("conn " + id + " no existent")
	}
	return x.(*pb.Connection), nil
}
func (mgr *connectManager) Remove(id string) error {
	_, ok := mgr.connMap.LoadAndDelete(id)
	if !ok {
		return errors.New("conn " + id + " no existent")
	}

	return nil
}
func (mgr *connectManager) AddOrUpData(conn *pb.Connection) {
	mgr.connMap.Store(conn.GetClientId(), conn)
}
func (mgr *connectManager) Oauth(id, token string) bool {
	c, err := mgr.Get(id)
	if err != nil {
		return false
	}
	if c.GetToken() != token {
		return false
	}
	return true
}

func (mgr *connectManager) AddWatcher(id string, ser pb.Stream_WatchStreamServer) {
	mgr.watcherMap.Store(id, ser)
}

func (mgr *connectManager) SendWatcher(id string, msg *pb.Message) error {
	c, ok := mgr.watcherMap.Load(id)
	if !ok {
		return errors.New("监听者不存在")
	}
	return c.(pb.Stream_WatchStreamServer).Send(msg)
}

func (mgr *connectManager) RemoveWatcher(id string) {
	mgr.watcherMap.Delete(id)
}
