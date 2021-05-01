package model

import (
	_ "github.com/go-sql-driver/mysql"

	"github.com/hsyodyssey/agoge/agogedb"
	"github.com/hsyodyssey/agoge/dao"
)

func GetAllTests(testtable *[]dao.Testtable) (err error) {
	if err = agogedb.MysqlDB.Find(testtable).Error; err != nil {
		return err
	}
	return nil
}

func GetFirstTest(testtable *dao.Testtable) (err error) {
	if err = agogedb.MysqlDB.First(testtable).Error; err != nil {
		return err
	}
	return nil
}
