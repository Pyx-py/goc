package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"goc/config"
	"goc/constant"
	"goc/models"
	"goc/pkg/db"
	"testing"
)

func TestRun(t *testing.T) {
	config.InitConfig("/home/pengyuxuan/go_pro/T-blog/goc/config/config_file/config.ini")
	conf := config.GetConfig()
	fmt.Printf("confObj:%+v\n", conf)
	mysqlUrl := fmt.Sprintf(constant.MysqlUrlStr, conf.Database.User, conf.Database.Password, conf.Database.Host, conf.Database.Port, conf.Database.DbName)
	db.InitMysql(mysqlUrl)
	var article = models.Article{}
	dbO := db.GetDB()
	defer dbO.Close()
	dbO = dbO.Table(models.Article{}.TableName())
	err := dbO.Where("id = ?", 3).Where("modified_user = ?", "李三").First(&article).Error
	if err == gorm.ErrRecordNotFound {
		fmt.Println("未查到记录")
		return
	}
	if err != nil {
		fmt.Println("查询记录出错")
		return
	}
	fmt.Println("查询记录成功")
	return
}

func InitSomeData(db *gorm.DB) {
	db = db.Table(models.Article{}.TableName())
	articleObj := models.Article{
		BaseModel:    models.BaseModel{},
		Title:        "测试数据文章标题",
		Content:      "测试数据文章内容",
		ModifiedUser: "李三",
		Like:         4,
	}
	articleObj1 := models.Article{
		BaseModel:    models.BaseModel{},
		Title:        "测试数据文章标题1",
		Content:      "测试数据文章内容1",
		ModifiedUser: "李四",
		Like:         3,
	}
	if err := db.Create(&articleObj).Error; err != nil {
		fmt.Printf("创建文章数据失败：%s", err.Error())
		panic("")
	}
	if err := db.Create(&articleObj1).Error; err != nil {
		fmt.Printf("创建文章数据失败：%s", err.Error())
		panic("")
	}

	db = db.Table(models.User{}.TableName())
	userObj := models.User{
		BaseModel: models.BaseModel{},
		UserName:  "李三",
		Password:  "admin",
	}
	if err := db.Create(&userObj).Error; err != nil {
		fmt.Printf("创建用户数据失败：%s", err.Error())
		panic("")
	}
}

func TestInitSomeData(t *testing.T) {
	config.InitConfig("/home/pengyuxuan/go_pro/T-blog/goc/config/config_file/config.ini")
	conf := config.GetConfig()
	fmt.Printf("confObj:%+v\n", conf)
	mysqlUrl := fmt.Sprintf(constant.MysqlUrlStr, conf.Database.User, conf.Database.Password, conf.Database.Host, conf.Database.Port, conf.Database.DbName)
	db.InitMysql(mysqlUrl)
	dbObj := db.GetDB()
	InitSomeData(dbObj)
	_ = dbObj.Close()
}

func TestChangeData(t *testing.T) {
	config.InitConfig("/home/pengyuxuan/go_pro/T-blog/goc/config/config_file/config.ini")
	conf := config.GetConfig()
	fmt.Printf("confObj:%+v\n", conf)
	mysqlUrl := fmt.Sprintf(constant.MysqlUrlStr, conf.Database.User, conf.Database.Password, conf.Database.Host, conf.Database.Port, conf.Database.DbName)
	db.InitMysql(mysqlUrl)
	var ar = models.Article{}
	dbObj := db.GetDB()
	defer dbObj.Close()
	dbObj = dbObj.Table(models.Article{}.TableName())
	dbObj.Where("id= ?", 1).First(&ar)
	err := dbObj.Delete(&ar).Error
	if err != nil {
		fmt.Println("删除记录出错")
		return
	}
	fmt.Println("删除记录成功")
}
