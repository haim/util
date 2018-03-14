package util

import "sync"

type Counter interface {
	Count() uint64
}

func NewCounter() Counter {
	return &defaultCounter{}
}

type defaultCounter struct {
	num     uint64
	numLock sync.RWMutex
}

func (c *defaultCounter) Count() uint64 {
	c.numLock.Lock()
	defer c.numLock.Unlock()
	id := c.num
	c.num++
	return id
}