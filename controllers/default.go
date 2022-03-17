package controllers

import (
	"fmt"
	"study/models"
	"time"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	log := logs.NewLogger()
	log.Debug("Geaaaa224445555")

	now := time.Now()
	year := now.Year()     //年
	month := now.Month()   //月
	day := now.Day()       //日
	hour := now.Hour()     //小时
	minute := now.Minute() //分钟
	second := now.Second() //秒
	date := fmt.Sprintf("%+v-%v-%+v %v:%v:%v\n", year, month, day, hour, minute, second)

	c.Ctx.WriteString(date)
	c.Data["Website"] = "beego.222me2223333"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"

}

func (c *MainController) Producer() {
	log := logs.NewLogger()

	p := models.Kafkap
	topic := "test"
	p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte("33333"),
	}, nil)
	// p.Flush(100)

	log.Debug("rrrrrrrrrrrrrrr")

	c.Data["Website"] = "beego.222me2223333"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"

}
