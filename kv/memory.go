package kv

import (
	"sync"
)

func NewMemoryKv() Kv {
	return &memoryKv{}
}

type memoryKv struct {
	values sync.Map

	//values atomic.Value
}

func (m *memoryKv) Initialization() {

}

func (m *memoryKv) Get(key interface{}) (interface{}, bool) {
	return m.values.Load(key)
}

func (m *memoryKv) Set(key interface{}, value interface{}) {
	m.values.Store(key, value)
}

func (m *memoryKv) Remove(key interface{}) {
	m.values.Delete(key)
}

func (m *memoryKv) IsExist(key interface{}) bool {
	_, b := m.values.Load(key)
	return b
}

//
//func (m *memoryKv) Lock(key interface{}) {
//
//
//	ch,ok:=m.lockMap[key]
//	if ok{
//		<-ch.(chan struct{})
//	}
//
//	//log.Println("!ok")
//	ch = make(chan struct{},1)
//	//<-ch.(chan struct{})
//	m.lockMap.LoadOrStore(key,ch)
//
//}
//
//func (m *memoryKv) UnLock(key interface{}) {
//	ch,ok:=m.lockMap.LoadAndDelete(key)
//	if !ok{
//		return
//	}
//	ch.(chan struct{})<- struct {}{}
//	close(ch.(chan struct{}))
//
//}
