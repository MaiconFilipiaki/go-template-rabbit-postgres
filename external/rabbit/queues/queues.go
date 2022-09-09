package queues

import (
	"github.com/streadway/amqp"
	"log"
)

type Queue string

var Queues = map[string]Queue{
	"GREEN-QUEUE": "green-queue",
	"RED-QUEUE":   "red-queue",
}

func Declare(ch *amqp.Channel, nameQueue string) (q amqp.Queue) {
	q, err := ch.QueueDeclare(
		nameQueue,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}
	return
}
