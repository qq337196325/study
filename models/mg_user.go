package models

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

type User struct {
	Name  string `json:"name"`  //NAMELOOKUP_TIME
	Email string `json:"email"` //NAMELOOKUP_TIME
}

//添加一条记录
func AddUser() {
	//新增一条数据
	UserRes := &User{Name: "小明", Email: "1234566@qq.com"}
	Mdb.User.InsertOne(context.TODO(), UserRes)
}

//删除一条记录
func DeleteOneUser() {
	// filter := bson.D{{"email", "1234566@qq.com"}}
	filter := bson.M{"email": "1234566@qq.com"}
	_, err := Mdb.User.DeleteOne(context.TODO(), filter)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}
}
