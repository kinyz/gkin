package connect

import (
	"errors"
	"gkin/pb"
	"sync"
)

type Manager interface {
	Get(id string) (*pb.Connection, error)
	Remove(id string) error
	AddOrUpData(conn *pb.Connection)
	Oauth(id, token string) bool
}

type manager struct {
	connects map[string]*pb.Connection
	lock     sync.RWMutex
}

func NewManager() Manager {
	return &manager{connects: make(map[string]*pb.Connection)}
}
func (mgr *manager) Get(id string) (*pb.Connection, error) {
	mgr.lock.RLock()
	defer mgr.lock.RUnlock()
	conn, ok := mgr.connects[id]
	if !ok {
		return nil, errors.New("conn " + id + " no existent")
	}
	return conn, nil
}

func (mgr *manager) Remove(id string) error {
	_, err := mgr.Get(id)
	if err != nil {
		return err
	}
	mgr.lock.Lock()
	defer mgr.lock.Unlock()
	//conn.Close()
	delete(mgr.connects, id)
	return nil
}

func (mgr *manager) AddOrUpData(conn *pb.Connection) {
	mgr.lock.Lock()
	mgr.connects[conn.GetClientId()] = conn
	mgr.lock.Unlock()
}

func (mgr *manager) Oauth(id, token string) bool {
	c, err := mgr.Get(id)
	if err != nil {
		return false
	}
	if c.GetToken() != token {
		return false
	}
	return true
}
