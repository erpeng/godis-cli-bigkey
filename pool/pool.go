package pool

//A Ele is a Item
type element struct {
	expireTime uint64
	lru        uint64
	lfu        uint64
	key        string
	valueType  string
	valueSize  uint64
}

var pool []*element

func initLen(l uint64) {
	pool = make([]*element, l)
}

func insert(e *element) {
	len := len(pool)
	//empty
	if pool[0] == nil {
		pool[0] = e
		return
		//full and e is least
	} else if pool[len-1] != nil && pool[0].valueSize > e.valueSize {
		return
		//full and e is biggest
	} else if pool[len-1] != nil && pool[len-1].valueSize < e.valueSize {
		//pool[0 : len-1] = pool[1:len]
		pool[len-1] = e
		return
	}

	//find the position
	pos := find(e, pool)
	if pos == 0 && pool[len-1] == nil {
		//pool[1:len] = pool[0 : len-1]
		pool[0] = e
		return
	} else {

	}

}

func find(e *element, pool []*element) (pos uint64) {
	for i, ele := range pool {
		if e.valueSize < ele.valueSize {
			pos = uint64(i)
			break
		}
	}
	// if i == len(pool)-1 {
	// 	return i
	// }
	return
}
