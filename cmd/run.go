package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/pflag"
	"goc/config"
	"goc/constant"
	"goc/pkg/db"
)

var confFile = pflag.StringP("config", "c", "", "config file path")

func main() {
	pflag.Parse()
	config.InitConfig(*confFile)
	conf := config.GetConfig()
	mysqlUrl := fmt.Sprintf(constant.MysqlUrlStr, conf.Database.User, conf.Database.Password, conf.Database.Host, conf.Database.Port, conf.Database.DbName)
	db.InitMysql(mysqlUrl)

}
