package controllers

import (
	"context"
	"study/models"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MongoController struct {
	beego.Controller
}

func (c *MongoController) Install() {
	log := logs.NewLogger()
	log.Debug("iuuuuuuuu")

	Mdb := models.Mdb
	UserRes := &models.User{Name: "小明", Email: "1234566@qq.com"}
	Mdb.User.InsertOne(context.TODO(), UserRes)

	c.Ctx.WriteString("3333444")
}

func (c *MongoController) Update() {
	c.Ctx.WriteString("Update")

	//根据ID更新
	objectId, _ := primitive.ObjectIDFromHex("622c8c7fb8edf83bdd465a03")
	filter := bson.M{"_id": objectId}

	//更新内容
	update := bson.M{
		"$set": bson.M{
			"email": "9999999@163.com",
		},
	}
	Mdb := models.Mdb
	_, err := Mdb.User.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		c.Ctx.WriteString("更新失败")
	}
}
