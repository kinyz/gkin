package pool

import (
	"gkin/pb"
	"sync"
)

var connectPool = sync.Pool{New: func() interface{} {
	return &pb.Connection{
		ClientId: "",

		Token: "",
		Key:   "",
	}
}}

func GetConnect() *pb.Connection {
	c := connectPool.Get().(*pb.Connection)
	c.ClientId = ""
	c.Token = ""
	c.Key = ""
	return c
}

func PutConnect(conn *pb.Connection) {
	connectPool.Put(conn)
}
