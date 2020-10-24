package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/hsyodyssey/agoge/config"

	"github.com/hsyodyssey/agoge/dao"
)

func GetAllTests(testtable *[]dao.Testtable) (err error) {
	if err = config.DB.Find(testtable).Error; err != nil {
		return err
	}
	return nil
}

func GetFirstTest(testtable *dao.Testtable) (err error) {
	if err = config.DB.First(testtable).Error; err != nil {
		return err
	}
	return nil
}
