package agogedb

import (
	"fmt"
	"log"

	"github.com/syndtr/goleveldb/leveldb"
)

func TestLevelDB() {
	db, err := leveldb.OpenFile("agogedb/test/level", nil)
	if err != nil {
		log.Fatal("Test LevelDb Go")
	}
	defer db.Close()
	iter := db.NewIterator(nil, nil)

	for iter.Next() {
		key := iter.Key()
		value := iter.Value()
		fmt.Printf("Key: %s | Value : %s\n", key, value)
	}

	iter.Release()
	err = iter.Error()
	if err != nil {
		log.Fatal("Test LevelDb Go")
	}
}

func initLevelDb() {
	db, err := leveldb.OpenFile("agogedb/test/level", nil)
	if err != nil {
		log.Fatal("Yikes!")
	}
	defer db.Close()
	err = db.Put([]byte("1"), []byte("1"), nil)
	err = db.Put([]byte("2"), []byte("2"), nil)
	err = db.Put([]byte("3"), []byte("3"), nil)
	db.Close()
}
