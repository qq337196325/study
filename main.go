package main

import (
	"study/models"
	_ "study/routers"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	orm.Debug = true

	//初始化Mongo与mysql
	MongodbObj := &models.Mongodb{}
	Mysql := &models.Mysql{}
	go MongodbObj.NewMongo()
	go Mysql.NewMysql()

	// go models.Consumers()
	// models.Producer()

	beego.Run()
}
