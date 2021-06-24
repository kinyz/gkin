package connect

import (
	"gkin/pb"
	"sync"
)

var Pool = sync.Pool{New: func() interface{} {
	return &pb.Connection{
		ClientId: "",

		Token: "",
		Key:   "",
	}
}}

func New() *pb.Connection {
	c := Pool.Get().(*pb.Connection)

	c.ClientId = ""
	c.Token = ""
	c.Key = ""
	return c
}
