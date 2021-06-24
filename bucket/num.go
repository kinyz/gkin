package bucket

func (b *Bucket) addConnNum() {
	b.connLock.Lock()
	b.nowConnNum++
	b.connLock.Unlock()
}
func (b *Bucket) reduceConnNum() {
	b.connLock.Lock()
	b.nowConnNum--
	b.connLock.Unlock()
}
func (b *Bucket) getConnNum() int {
	b.connLock.RLock()
	defer b.connLock.RUnlock()
	return b.nowConnNum
}

func (b *Bucket) getConnMaxNum() int {
	return b.maxConnNum
}
