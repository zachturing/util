package queue

//
// import (
// 	"fmt"
// 	amqp "github.com/rabbitmq/amqp091-go"
// 	"github.com/zachturing/util/log"
// )
//
// const (
// 	// 连接地址从控制台获取
// 	// http://49.234.241.169:15672
// 	host = "49.234.241.169"
// 	// 角色名称, 位于【角色管理】页面
// 	username = "admin"
// 	// 角色密钥, 位于【角色管理】页面
// 	password = "ulbJujmfGHMEW0oI"
// 	// 要使用vhost全称
// 	vhost = "test-vhost"
// )
//
// func failOnError(err error, msg string) {
// 	if err != nil {
// 		panic(fmt.Sprintf("%s: %s", msg, err))
// 	}
// }
//
// func main() {
// 	// 创建连接
// 	conn, err := amqp.Dial("amqp://" + username + ":" + password + "@" + host + ":15672/" + vhost)
// 	// conn, err := amqp.Dial("amqp://" + username + ":" + password + "@" + host + ":5672/")
// 	failOnError(err, "Failed to connect to RabbitMQ")
// 	defer func(conn *amqp.Connection) {
// 		err := conn.Close()
// 		if err != nil {
// 			log.Errorf("Failed to close connection: %s", err)
// 		}
// 	}(conn)
//
// 	// 建立通道
// 	ch, err := conn.Channel()
// 	failOnError(err, "Failed to open a channel")
// 	defer func(ch *amqp.Channel) {
// 		err := ch.Close()
// 		if err != nil {
// 			log.Errorf("Failed to close channel: %s", err)
// 		}
// 	}(ch)
//
// 	// 声明消息队列
// 	queue, err := ch.QueueDeclare(
// 		"aipaper-queue",
// 		false,
// 		false,
// 		false,
// 		false,
// 		nil)
// 	failOnError(err, "Failed to declare a queue")
//
// 	// 消息内容
// 	body := "Hello World!"
//
// 	// 发布消息到指定的消息队列中
// 	err = ch.Publish(
// 		"",         // exchange
// 		queue.Name, // routing key
// 		false,      // mandatory
// 		false,      // immediate
// 		amqp.Publishing{
// 			ContentType: "text/plain",
// 			Body:        []byte(body),
// 		})
// 	failOnError(err, "Failed to publish a message")
//
// 	log.Debugf("[Producer(Hello world)] Sent %s", body)
// }
