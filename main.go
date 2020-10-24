package main

import (
	"github.com/hsyodyssey/agoge/config"
	"github.com/hsyodyssey/agoge/dao"
	"github.com/hsyodyssey/agoge/router"

	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

var err error

func main() {
	config.DB, err = gorm.Open(mysql.Open(config.DbURL(config.BuildDBConfig())), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	config.DB.AutoMigrate(&dao.Testtable{})

	r := router.SetupRouter()
	r.Run()
}
