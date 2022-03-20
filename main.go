package main

import (
	"study/grpc"
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

	models.NewRedis()

	grpcServer := grpc.Server{}
	go grpcServer.NewGrpcServer()
	// 这里是test1分支
	// redis.Dial("tcp", "127.0.0.1:6379")

	go models.Consumers()
	models.Producer()

	beego.Run()
}
