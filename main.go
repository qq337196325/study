package main

import (
	"study/models"
	_ "study/routers"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	MongodbObj := &models.Mongodb{}
	go MongodbObj.NewMongo()

	//orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/beego?charset=utf8")

	// go models.Consumers()
	// models.Producer()

	beego.Run()
}
