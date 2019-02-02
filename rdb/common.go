package rdb

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

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
	for b, err := ReadBytes(f, 1); err == nil; b, err = ReadBytes(f, 1) {

		if b[0] == RDB_OPCODE_EXPIRETIME_MS {
			expire, _ := ReadBytes(f, rdbExpireTimeLen)
			expireInt := binary.LittleEndian.Uint64(expire)
			fmt.Println(expireInt)
		} else if b[0] == RDB_OPCODE_FREQ {
			lfu, _ := ReadBytes(f, rdbLfuLen)
			fmt.Println(binary.LittleEndian.Uint16(lfu))
		} else if b[0] == RDB_OPCODE_IDLE {

		} else if b[0] == RDB_OPCODE_AUX {
		} else if b[0] == RDB_OPCODE_EOF {
		} else if b[0] == RDB_OPCODE_EXPIRETIME {
		} else if b[0] == RDB_OPCODE_RESIZEDB {
		} else if b[0] == RDB_OPCODE_SELECTDB {
		}

		valueType := int(b[0])
		b, _ := ReadBytes(f, 1)
		len := readRdbLength(f, b[0])
		key := readKey(f, len)
		fmt.Println(key)
		b, _ = ReadBytes(f, 1)
		len = readRdbLength(f, b[0])
		m[valueType](f, len)

	}
}
