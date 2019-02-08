package rdb

import (
	"os"
)

const (
	RDB_OPCODE_IDLE = iota + 248
	RDB_OPCODE_FREQ
	RDB_OPCODE_AUX
	RDB_OPCODE_RESIZEDB
	RDB_OPCODE_EXPIRETIME_MS
	RDB_OPCODE_EXPIRETIME
	RDB_OPCODE_SELECTDB
	RDB_OPCODE_EOF
)

const (
	rdbExpireTimeLen   = 8
	rdbMagicVersionLen = 4
	rdbLfuLen          = 1
)

func readAux(f *os.File) {
	for i := 1; i < 3; i++ {
		lenFlag, _ := ReadBytes(f, 1)
		len, isInt, _ := readRdbLength(f, lenFlag[0])
		if isInt {
			Printf("%d ", len)
		} else {
			b, _ := ReadBytes(f, len)
			Printf("%s ", b)
		}
	}
}

func readDbNum(f *os.File) {
	lenFlag, _ := ReadBytes(f, 1)
	len, _, _ := readRdbLength(f, lenFlag[0])
	Printf("db:%d\n", len)
}

func readDbSize(f *os.File) {
	for i := 1; i < 3; i++ {
		lenFlag, _ := ReadBytes(f, 1)
		len, _, _ := readRdbLength(f, lenFlag[0])
		if i == 1 {
			Printf("db-size:%d\n", len)
		}
		if i == 2 {
			Printf("expire-size:%d\n", len)
		}
	}
}

func readEOF(f *os.File) {
	len, _ := ReadBytes(f, 8)
	Printf("EOF checksum:%v\n", len)
}

func readIdle(f *os.File) (lru uint64) {
	lenFlag, _ := ReadBytes(f, 1)
	len, _, _ := readRdbLength(f, lenFlag[0])
	Printf("%d\n", len)
	return len
}
