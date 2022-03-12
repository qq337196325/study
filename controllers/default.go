package controllers

import (
	"study/models"

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
