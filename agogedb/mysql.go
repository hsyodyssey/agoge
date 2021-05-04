package agogedb

import (
	"fmt"

	"github.com/hsyodyssey/agoge/dao"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MysqlDB *gorm.DB
var err error

// DBConfig represents db configuration
type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host:     "127.0.0.1",
		Port:     3306,
		User:     "hsylocal",
		Password: "hanslocal",
		DBName:   "agoge",
	}
	return &dbConfig
}

func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}

func InitMysql() {
	MysqlDB, err = gorm.Open(mysql.Open(DbURL(BuildDBConfig())), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	MysqlDB.AutoMigrate(&dao.Testtable{})

}
