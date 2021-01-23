package main

import (
	"fmt"
	"log"

	"github.com/syndtr/goleveldb/leveldb"
)

var err error

func main() {
	// config.DB, err = gorm.Open(mysql.Open(config.DbURL(config.BuildDBConfig())), &gorm.Config{})

	// if err != nil {
	// 	panic(err)
	// }

	// config.DB.AutoMigrate(&dao.Testtable{})

	// r := router.SetupRouter()
	// r.Run()

	db, err := leveldb.OpenFile("db/test", nil)
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
	db, err := leveldb.OpenFile("db/test", nil)
	if err != nil {
		log.Fatal("Yikes!")
	}
	defer db.Close()
	err = db.Put([]byte("1"), []byte("1"), nil)
	err = db.Put([]byte("2"), []byte("2"), nil)
	err = db.Put([]byte("3"), []byte("3"), nil)
	db.Close()
}
