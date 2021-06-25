package kv

type Kv interface {

	// Initialization 初始化
	Initialization()

	Get(key interface{}) (interface{}, bool)
	Set(key interface{}, value interface{})
	Remove(key interface{})

	//Lock(key interface{})
	//UnLock(key interface{})

	IsExist(key interface{}) bool

	//	GetMap()map[interface{}]interface{}
}
