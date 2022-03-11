package main

import (
	"fmt"
	"study/models"
	_ "study/routers"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	MongodbObj := &models.Mongodb{}
	MongodbObj.NewMongo()

	// models.AddUser()
	models.DeleteOneUser()
	// fmt.Println(err)
	fmt.Println("Connected to MongoDB!")
	beego.Run()
}
