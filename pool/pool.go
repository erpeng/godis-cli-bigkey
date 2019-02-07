package pool

import "fmt"

//Element save redis key/value
type Element struct {
	ExpireTime uint64
	Lru        uint64
	Lfu        uint16
	Key        string
	ValueType  int
	ValueSize  uint64
}

var pool []*Element
var poolLength int

//InitLen pool length
func InitLen(l int) {
	poolLength = l
	pool = make([]*Element, 0, l)
}

//Insert insert element
func Insert(e *Element) {
	len := len(pool)

	//find the position
	pos := find(e, pool)

	var poolTmp []*Element
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

func find(e *Element, pool []*Element) (pos int) {
	index := 0
	for i, ele := range pool {
		if e.ValueSize < ele.ValueSize {
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

//PrintPool print pool content
func PrintPool() {
	for _, ele := range pool {
		fmt.Printf("key:%s,valueSize:%d,valueType:%d", ele.Key, ele.ValueSize, ele.ValueType)
		fmt.Printf("expireTime:%d,lfu:%d,lru:%d\n", ele.ExpireTime, ele.Lfu, ele.Lru)
	}
}
