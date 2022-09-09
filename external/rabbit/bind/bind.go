package bind

import (
	"github.com/streadway/amqp"
	"log"
)

type RoutingKey string

var Routingkies = map[string]RoutingKey{
	"GREEN-ROUTING": "green-queue",
	"RED-ROUTING":   "red-queue",
}

func BindQueueRoutingExchange(ch *amqp.Channel, queueName string, routingName string, exName string) {
	err := ch.QueueBind(
		queueName,
		routingName,
		exName,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}
}
