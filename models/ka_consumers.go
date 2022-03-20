package models

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/beego/beego/v2/core/config"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var Kafkac *kafka.Consumer

// 消费者
func Consumers() {
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	kafka_ip, _ := config.String("kafka_ip")
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":     kafka_ip, //broker,
		"broker.address.family": "v4",
		"group.id":              "test-consumer-group", //group,
		// "session.timeout.ms":    6000,
		"auto.offset.reset": "earliest"})
	// defer c.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create consumer: %s\n", err)
		os.Exit(1)
	}

	Kafkac = c
	err = c.SubscribeTopics([]string{"test"}, nil)
	run := true

	for run {
		select {
		case sig := <-sigchan:
			fmt.Printf("Caught signal %v: terminating\n", sig)
			run = false
		default:
			ev := c.Poll(100)
			if ev == nil {
				continue
			}

			switch e := ev.(type) {
			case *kafka.Message:
				fmt.Printf("接受到消息： %s:%s\n",
					e.TopicPartition, string(e.Value))
				if e.Headers != nil {
					fmt.Printf("%% Headers: %v\n", e.Headers)
				}
			case kafka.Error:
				// Errors should generally be considered
				// informational, the client will try to
				// automatically recover.
				// But in this example we choose to terminate
				// the application if all brokers are down.
				fmt.Fprintf(os.Stderr, "%% Error: %v: %v\n", e.Code(), e)
				if e.Code() == kafka.ErrAllBrokersDown {
					run = false
				}
			default:
				fmt.Printf("Ignored %v\n", e)
			}
		}
	}
}
