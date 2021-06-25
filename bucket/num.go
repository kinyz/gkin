package bucket

import "sync/atomic"

func (b *Bucket) addConnNum() {

	atomic.AddInt64(&b.nowConnNum, 1)
	//b.connLock.Lock()
	//b.nowConnNum++
	//b.connLock.Unlock()
}
func (b *Bucket) reduceConnNum() {
	atomic.AddInt64(&b.nowConnNum, -1)
}
func (b *Bucket) getConnNum() int64 {
	return b.nowConnNum
}

func (b *Bucket) getConnMaxNum() int64 {
	return b.maxConnNum
}
