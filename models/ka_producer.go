package models

import (
	"fmt"
	"github.com/beego/beego/v2/core/config"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var Kafkap *kafka.Producer

// 生产者
func Producer() {
	kafka_ip, _ := config.String("kafka_ip")
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": kafka_ip})

	if err != nil {
		panic(err)
	}

	// Delivery report handler for produced messages
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("发送失败: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("发送消息： %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	// Produce messages to topic (asynchronously)
	// topic := "test"
	// for _, word := range []string{"pppppp", "oooooo", "iiiiii", "uuuu", "rrrrrr", "eeeee", "wwwww"} {
	// 	fmt.Printf("4445555%v \n", word)
	// 	p.Produce(&kafka.Message{
	// 		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
	// 		Value:          []byte(word),
	// 	}, nil)
	// }
	Kafkap = p

	// Wait for message deliveries before shutting down
	p.Flush(1) //刷新并等待未完成的消息和请求以完成传递

}
