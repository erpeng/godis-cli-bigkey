package rdb

const (
	RDB_MAGIC = "REDIS"

	RDB_6BITLEN  = 0
	RDB_14BITLEN = 1
	RDB_32BITLEN = 0x80
	RDB_64BITLEN = 0x81
	RDB_ENCVAL   = 3

	RDB_ENC_INT8  = 0
	RDB_ENC_INT16 = 1
	RDB_ENC_INT32 = 2
	RDB_ENC_LZF   = 3

	rdb6bitlen  = 1
	rdb14bitlen = 2
	rdb32bitlen = 5
	rdb64bitlen = 9

	rdbenc8len  = 2
	rdbenc16len = 3
	rdbenc32len = 5
)

const (
	RDB_TYPE_STRING = iota
	RDB_TYPE_LIST
	RDB_TYPE_SET
	RDB_TYPE_ZSET
	RDB_TYPE_HASH
	RDB_TYPE_ZSET_2
	RDB_TYPE_MODULE
	RDB_TYPE_MODULE_2
	_
	RDB_TYPE_HASH_ZIPMAP
	RDB_TYPE_LIST_ZIPLIST
	RDB_TYPE_SET_INTSET
	RDB_TYPE_ZSET_ZIPLIST
	RDB_TYPE_HASH_ZIPLIST
	RDB_TYPE_LIST_QUICKLIST
	RDB_TYPE_STREAM_LISTPACKS
)

func readRdbLength(b byte) (t int, l int64) {
	flag := (int(b) & 0xC0) >> 6
	if flag == RDB_6BITLEN {

	} else if flag == RDB_14BITLEN {

	} else if flag == RDB_ENCVAL {

	} else if b == RDB_32BITLEN {

	} else if b == RDB_64BITLEN {

	} else {
		panic("Unknown len")
	}
	return

}
