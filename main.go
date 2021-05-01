package main

import (
	"github.com/hsyodyssey/agoge/agogedb"
)

func main() {
	agogedb.TestLevelDB()
	agogedb.TestRocksDB()
}
