package controllers

import (
	"context"
	"fmt"
	"study/models"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoController struct {
	beego.Controller
}

func (c *MongoController) Install() {
	log := logs.NewLogger()

	Mdb := models.Mdb
	UserRes := &models.User{Name: "小明", Email: "1234566@qq.com"}
	_, err := Mdb.User.InsertOne(context.TODO(), UserRes)
	if err != nil {
		log.Debug("新增失败")
		c.Ctx.WriteString("新增失败")
		return
	}

	c.Ctx.WriteString("新增一条数据成功\n")

	//批量插入
	// trainers := []interface{}{&models.User{Name: "小明1", Email: "1234566@qq.com"}, &models.User{Name: "小明2", Email: "1234566@qq.com"}}
	trainers := []interface{}{}
	trainers = append(trainers, &models.User{Name: "小明3", Email: "1234566@qq.com"})
	trainers = append(trainers, &models.User{Name: "小明4", Email: "1234566@qq.com"})

	_, err = Mdb.User.InsertMany(context.TODO(), trainers)
	if err != nil {
		log.Debug("批量新增失败")
		c.Ctx.WriteString("批量新增失败")
		return
	}
	c.Ctx.WriteString("批量新增数据成功\n")
}

func (c *MongoController) Update() {
	c.Ctx.WriteString("Update")
	Mdb := models.Mdb

	//根据ID更新
	objectId, _ := primitive.ObjectIDFromHex("622c8c7fb8edf83bdd465a03")
	filter := bson.M{"_id": objectId}

	//更新内容
	update := bson.M{
		"$set": bson.M{
			"email": "9999999@163.com",
		},
	}

	_, err := Mdb.User.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		c.Ctx.WriteString("更新一条数据失败\n")
	}
	c.Ctx.WriteString("更新一条数据成功\n")

	//批量更新
	filter = bson.M{"name": "小明3"}
	//更新内容
	update = bson.M{
		"$set": bson.M{
			"email": "aaaaaaa@163.com",
		},
	}
	_, err = Mdb.User.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		c.Ctx.WriteString("更新一条数据失败\n")
	}
	c.Ctx.WriteString("更新一条数据成功\n")
}

func (c *MongoController) Find() {
	Mdb := models.Mdb

	//根据ID查找
	var user models.User
	// objectId, _ := primitive.ObjectIDFromHex("622c8c7fb8edf83bdd465a03")
	filter := bson.M{}
	err := Mdb.User.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		c.Ctx.WriteString("查找一条数据失败\n")
		// return
	}
	fmt.Printf("查询到得单条数据: %+v\n", user)

	var user2 models.User
	project := bson.M{"name": true}
	projection := options.FindOne().SetProjection(project)
	Mdb.User.FindOne(context.TODO(), bson.D{{}}, projection).Decode(&user2)
	fmt.Printf("查询到得单条数据2: %+v\n", user2)
	fmt.Printf("3333: %+v\n", projection)
	c.Ctx.WriteString("查询到得单条数据\n")

	//批量查找
	findOptions := options.Find()
	findOptions.SetLimit(5)
	cur, err := Mdb.User.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		c.Ctx.WriteString("查找多条数据失败\n")
		return
	}

	//循环取出
	var users []*models.User
	for cur.Next(context.TODO()) {
		// create a value into which the single document can be decoded
		var u models.User
		err := cur.Decode(&u)
		if err != nil {
			c.Ctx.WriteString("取出数据失败\n")
			return
		}

		users = append(users, &u)
	}
	fmt.Printf("查询到得批量数据: %+v\n", *users[1])
	c.Ctx.WriteString("查询到得批量数据\n")

}
