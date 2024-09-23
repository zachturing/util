package main

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/zachturing/util/log"
)

const (
	// 连接地址从控制台获取
	host = "172.81.233.2"
	// 角色名称, 位于【角色管理】页面
	username = "admin"
	// 角色密钥, 位于【角色管理】页面
	password = "ulbJujmfGHMEW0oI"
	// 要使用vhost全称
	vhost = "test-vhost"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Error("%s: %s", msg, err)
	}
}

func main() {
	// 创建连接
	conn, err := amqp.Dial("amqp://" + username + ":" + password + "@" + host + ":5672/" + vhost)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer func(conn *amqp.Connection) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	// 建立通道
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer func(ch *amqp.Channel) {
		err := ch.Close()
		if err != nil {
			log.Errorf("Failed to close channel: %s", err)
		}
	}(ch)

	// 创建消费者并消费指定消息队列中的消息
	msgs, err := ch.Consume(
		"test", // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	// 持续获取消息队列中的消息
	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Debugf("Received a message: %s", d.Body)
		}
	}()

	log.Debugf(" [Consumer(Hello world)] Waiting for messages.")
	<-forever
}
