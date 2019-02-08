package rdb

import (
	"bytes"
	"encoding/binary"
	"io"
	"os"

	"github.com/erpeng/godis-cli-bigkey/pool"
)

//DEBUG open debug mode
var DEBUG bool

//ReadBytes read n bytes
func ReadBytes(r io.Reader, n uint64) ([]byte, error) {
	b := make([]byte, n)
	_, err := io.ReadFull(r, b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

//EqualBytes compare two bytes
func EqualBytes(b []byte, b1 []byte) bool {
	return bytes.Equal(b, b1)
}

//Load load a rdb file
func Load(f *os.File) {
	var expireInt uint64
	var lfuInt uint16
	var lruInt uint64
	for b, err := ReadBytes(f, 1); err == nil; b, err = ReadBytes(f, 1) {

		if b[0] == RDB_OPCODE_EXPIRETIME_MS {
			expire, _ := ReadBytes(f, rdbExpireTimeLen)
			expireInt = binary.LittleEndian.Uint64(expire)
			Println(expireInt)
		} else if b[0] == RDB_OPCODE_FREQ {
			lfu, _ := ReadBytes(f, rdbLfuLen)
			lfuInt = binary.LittleEndian.Uint16(lfu)
			Println(lfuInt)
		} else if b[0] == RDB_OPCODE_IDLE {
			lruInt = readIdle(f)
		} else if b[0] == RDB_OPCODE_AUX {
			readAux(f)
		} else if b[0] == RDB_OPCODE_EOF {
			readEOF(f)
		} else if b[0] == RDB_OPCODE_EXPIRETIME {
			//new rdb version don't use this type
		} else if b[0] == RDB_OPCODE_RESIZEDB {
			readDbSize(f)
		} else if b[0] == RDB_OPCODE_SELECTDB {
			readDbNum(f)
		} else {
			valueType := int(b[0])
			Printf("valueType:%d\n", valueType)
			b, _ := ReadBytes(f, 1)
			len, _, _ := readRdbLength(f, b[0])
			key := readKey(f, len)
			Printf("key:%s\n", key)
			valueLen := readValue(f, valueType)
			ele := &pool.Element{ValueType: valueType, Key: key, ValueSize: valueLen,
				ExpireTime: expireInt, Lfu: lfuInt, Lru: lruInt}
			pool.Insert(ele)
			expireInt = 0
			lfuInt = 0
			lruInt = 0
		}
	}
}

func readValue(f *os.File, valueType int) (len uint64) {
	var length uint64
	m[valueType](f, &length)
	return length
}
