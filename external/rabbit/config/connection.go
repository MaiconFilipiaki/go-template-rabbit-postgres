package config

import (
	"github.com/streadway/amqp"
	"log"
)

func createConnectionRabbit() (conn *amqp.Connection) {
	conn, err := amqp.Dial("amqp://root:root@localhost:5672/")
	if err != nil {
		log.Fatal(err)
	}
	return
}

func CreateChannel() (ch *amqp.Channel) {
	conn := createConnectionRabbit()
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	//defer ch.Close()
	return
}

func SendMsgRabbit(ch *amqp.Channel, ex string, msg string, route string) {
	err := ch.Publish(
		ex,
		route,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		})
	if err != nil {
		log.Fatal(err)
	}
}
