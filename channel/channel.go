package channel

import (
	"context"
	"errors"
	"sync"
)

type Channel struct {
	ctx     context.Context
	cancel  context.CancelFunc
	c       chan interface{}
	isClose bool
	lock    sync.RWMutex
}

func NewChannel(cap int) *Channel {
	ctx, cancel := context.WithCancel(context.Background())
	return &Channel{ctx: ctx, cancel: cancel, isClose: false, c: make(chan interface{}, cap)}
}

func (c *Channel) Close() {
	c.lock.Lock()
	c.cancel()
	close(c.c)
	c.isClose = true
	c.lock.Unlock()
}

func (c *Channel) Done() <-chan struct{} {
	return c.ctx.Done()
}

func (c *Channel) Send(x interface{}) error {
	c.lock.RLock()
	defer c.lock.RUnlock()
	if c.isClose {
		return errors.New("通道已关闭")
	}
	c.c <- x
	return nil
}

func (c *Channel) Get() (chan interface{}, error) {
	c.lock.RLock()
	defer c.lock.RUnlock()
	if c.isClose {
		return nil, errors.New("通道已关闭")
	}
	return c.c, nil
}

func (c *Channel) Cap() int {
	return cap(c.c)
}
