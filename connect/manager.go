package connect

import (
	"errors"
	"sync"
)

type Manager interface {
	Get(id string) (Connect, error)
	Remove(id string) error
	AddOrUpData(conn Connect)
}

type manager struct {
	connects map[string]Connect
	lock     sync.RWMutex
}

func NewManager() Manager {
	return &manager{connects: make(map[string]Connect)}
}
func (mgr *manager) Get(id string) (Connect, error) {
	mgr.lock.RLock()
	defer mgr.lock.RUnlock()
	conn, ok := mgr.connects[id]
	if !ok {
		return nil, errors.New("conn " + id + " no existent")
	}
	return conn, nil
}

func (mgr *manager) Remove(id string) error {
	conn, err := mgr.Get(id)
	if err != nil {
		return err
	}
	mgr.lock.Lock()
	defer mgr.lock.Unlock()
	conn.Close()
	delete(mgr.connects, id)
	return nil
}

func (mgr *manager) AddOrUpData(conn Connect) {
	mgr.lock.Lock()
	mgr.connects[conn.GetClientId()] = conn
	mgr.lock.Unlock()
}
