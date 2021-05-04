package agogedb

import "fmt"

func NewDB(s string) {
	switch s {
	case "leveldb":
		fmt.Println("levelDB")
	case "rocksdb":
		fmt.Println("rocksDB")
	case "mysql":
		fmt.Println("mysql")
	}
}
