package agogedb

import (
	"fmt"
	"log"

	"github.com/tecbot/gorocksdb"
)

func TestRocksDB() {
	bbto := gorocksdb.NewDefaultBlockBasedTableOptions()
	bbto.SetBlockCache(gorocksdb.NewLRUCache(3 << 30))
	opts := gorocksdb.NewDefaultOptions()
	opts.SetBlockBasedTableFactory(bbto)
	opts.SetCreateIfMissing(true)
	db, err := gorocksdb.OpenDb(opts, "agogedb/test/rocks/test")

	if err != nil {
		log.Fatal("Something err happened in Rocks")
	}

	wo := gorocksdb.NewDefaultWriteOptions()
	ro := gorocksdb.NewDefaultReadOptions()
	err = db.Put(wo, []byte("foo"), []byte("bar"))
	err = db.Put(wo, []byte("1"), []byte("11"))
	err = db.Put(wo, []byte("2"), []byte("22"))

	err = db.Put(wo, []byte("3"), []byte("33"))
	it := db.NewIterator(ro)
	defer it.Close()
	it.Seek([]byte("1"))

	for it = it; it.Valid(); it.Next() {
		key := it.Key()
		value := it.Value()
		fmt.Printf("[Rocksdb] Key: %s Value: %s\n", key.Data(), value.Data())
		key.Free()
		value.Free()

	}

	if err := it.Err(); err != nil {
		log.Fatal(err)
	}

}
