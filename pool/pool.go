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
var poolLength int

func initLen(l int) {
	poolLength = l
	pool = make([]*element, 0, l)
}

func insert(e *element) {
	len := len(pool)

	//find the position
	pos := find(e, pool)

	var poolTmp []*element
	//empty
	if len == 0 {
		pool = append(pool, e)
		return
		//full and e is least
	} else if len == poolLength && pos == 0 {
		return
		//full and e is biggest
	} else if len == poolLength && pos == len {
		//pool[0 : len-1] = pool[1:len]
		poolTmp = append(poolTmp, pool[1:len]...)
		poolTmp = append(poolTmp, e)
		pool = poolTmp
		return
		//full ane e is middle
	} else if len == poolLength {

		poolTmp = append(poolTmp, pool[1:pos]...)
		poolTmp = append(poolTmp, e)
		poolTmp = append(poolTmp, pool[pos:len]...)
		pool = poolTmp
		//not full
	} else if len != poolLength && pos == 0 {
		poolTmp = append(poolTmp, e)
		poolTmp = append(poolTmp, pool...)

		pool = poolTmp
		//not full
	} else if len != poolLength && pos == len {
		pool = append(pool, e)
	} else if len != poolLength {
		poolTmp = append(poolTmp, pool[0:pos]...)
		poolTmp = append(poolTmp, e)
		poolTmp = append(poolTmp, pool[pos:len]...)
		pool = poolTmp
	}

}

func find(e *element, pool []*element) (pos int) {
	index := 0
	for i, ele := range pool {
		if e.valueSize < ele.valueSize {
			pos = i
			break
		}
		index++
	}
	if index == len(pool) {
		return len(pool)
	}
	return pos
}
