package main

import "github.com/rabbitmq/amqp091-go"

func demo() {
	conn, err := amqp091.Dial("")
	if err != nil {
		return
	}

	_ = conn
}
