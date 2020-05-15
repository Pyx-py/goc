package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"sync"
)

var one sync.Once
var db *gorm.DB

func InitMysql(mysqlUrl string) {
	one.Do(func() {
		dbMysql, err := gorm.Open("mysql", mysqlUrl)
		fmt.Printf("mysqlUrl:%s\n", mysqlUrl)
		if err != nil {
			fmt.Printf("连接数据库失败%s", err.Error())
			panic("connect failed")
		}
		// 关闭复数表名
		dbMysql.SingularTable(true)
		// 开启日志打印
		dbMysql.LogMode(true)
		// 关闭全表更新
		dbMysql.BlockGlobalUpdate(true)

		db = dbMysql
	})
}
func GetDB() *gorm.DB {
	return db
}
