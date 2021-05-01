package main

import (
	"fmt"
	"time"

	"github.com/hsyodyssey/agoge/agogedb"
)

func main() {
	// agogedb.TestLevelDB()

	db := agogedb.InitLevelDB()
	iter := db.NewIterator(nil, nil)

	for iter.Next() {
		key := iter.Key()
		value := iter.Value()
		t := time.Now()
		fmt.Printf("[LevelDB] [%d-%02d-%02d %02d:%02d:%02d] Key: %s | Value : %s\n", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), key, value)
	}

	iter.Release()
	db.Close()
	// agogedb.TestRocksDB()
}
