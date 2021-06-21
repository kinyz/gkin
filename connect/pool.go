package connect

import (
	"sync"
)

var Pool = sync.Pool{New: func() interface{} {
	return &Connection{
		ClientId: "",
		Oauth:    false,
		Token:    "",
		Key:      "",
	}
}}

func New() Connect {
	c := Pool.Get().(*Connection)

	c.Oauth = false
	c.ClientId = ""
	c.Token = ""
	c.Key = ""
	return c
}
