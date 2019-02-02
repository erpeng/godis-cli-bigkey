package main

import (
	"berkshire/rdb"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("./rdb.rdb")
	if err != nil {
		fmt.Printf("Open rdb file error:%v", err)
		os.Exit(1)
	}
	b, err := rdb.ReadBytes(f, 5)
	if err != nil {
		fmt.Printf("Read Rdb magic  error:%v", err)
		os.Exit(1)
	}
	if !rdb.EqualBytes(b, []byte(rdb.RDB_MAGIC)) {
		fmt.Printf("Error:%s\n", "invalid rdb file ")
		os.Exit(1)
	}
	b, err = rdb.ReadBytes(f, 4)
	if err != nil {
		fmt.Printf("Read Rdb version  error:%v", err)
		os.Exit(1)
	}
	fmt.Printf("\rRdb Version:%s\n", string(b))
	rdb.Load(f)
}
