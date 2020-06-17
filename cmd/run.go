package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"goc/config"
	"goc/constant"
	"goc/pkg/db"
	"goc/pkg/middlerwares"
	"goc/router"
)

var confFile = pflag.StringP("config", "c", "", "config file path")

func main() {
	log.SetReportCaller(true)
	pflag.Parse()
	config.InitConfig(*confFile)
	conf := config.GetConfig()
	mysqlUrl := fmt.Sprintf(constant.MysqlUrlStr, conf.Database.User, conf.Database.Password, conf.Database.Host, conf.Database.Port, conf.Database.DbName)
	db.InitMysql(mysqlUrl)
	log.Info("database ready")
	r := gin.Default()
	r.Use(middlerwares.CORSMiddleware())
	router.Router(r)
	if err := r.Run(conf.Server.Port); err != nil {
		log.Error("blog-server start failed")
		panic(err)
	}
	log.Info("blog-server START SUCCESS")
}
