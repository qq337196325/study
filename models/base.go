package models

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Mdb *Mongodb

type Mongodb struct {
	Client *mongo.Client
	User   *mongo.Collection
}

func (db *Mongodb) NewMongo() {
	param := fmt.Sprintf("mongodb://127.0.0.1:27017")
	clientOptions := options.Client().ApplyURI(param)

	// 建立客户端连接
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
		return
	}

	// 检查连接情况
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
		return
	}

	//指定要操作的数据集
	db.User = client.Database("testmongo").Collection("user")
	db.Client = client
	Mdb = db
}
