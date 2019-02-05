package rdb

import (
	"fmt"
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
		len, isInt := readRdbLength(f, lenFlag[0])
		if isInt {
			fmt.Printf("%d ", len)
		} else {
			b, _ := ReadBytes(f, len)
			fmt.Printf("%s ", b)
		}
	}
}

func readDbNum(f *os.File) {
	lenFlag, _ := ReadBytes(f, 1)
	len, _ := readRdbLength(f, lenFlag[0])
	fmt.Printf("db:%d\n", len)
}

func readDbSize(f *os.File) {
	for i := 1; i < 3; i++ {
		lenFlag, _ := ReadBytes(f, 1)
		len, _ := readRdbLength(f, lenFlag[0])
		if i == 1 {
			fmt.Printf("db-size:%d\n", len)
		}
		if i == 2 {
			fmt.Printf("expire-size:%d\n", len)
		}
	}
}

func readEOF(f *os.File) {
	len, _ := ReadBytes(f, 1)
	fmt.Printf("EOF:%d\n", len[0])
}

func readIdle(f *os.File) {
	lenFlag, _ := ReadBytes(f, 1)
	len, _ := readRdbLength(f, lenFlag[0])
	fmt.Printf("%d\n", len)
}
