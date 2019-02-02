package heap

//A Ele is a Item
type ele struct {
	expireTime uint64
	lru        uint64
	lfu        uint64
	key        string
	valueType  string
	valueSize  uint64
}
