package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/erpeng/godis-cli-bigkey/pool"
	"github.com/erpeng/godis-cli-bigkey/rdb"
)

func main() {
	debug := flag.Bool("debug", false, "open debug mode")
	topN := flag.Int("topn", 100, "output topn keys")
	flag.Parse()
	rdb.DEBUG = *debug
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
	pool.InitLen(*topN)
	rdb.Load(f)
	pool.PrintPool()
}
