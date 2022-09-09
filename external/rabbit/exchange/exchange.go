package exchange

import (
	"github.com/streadway/amqp"
	"log"
)

type KindExchange string

var KindExchanges = map[string]KindExchange{
	"FANOUT": "fanout",
	"DIRECT": "direct",
}

type Exchange string

var Exchanges = map[string]Exchange{
	"GREEN": "green",
	"RED":   "red",
}

func Declare(ch *amqp.Channel, name string, kind string) {
	err := ch.ExchangeDeclare(
		name,
		kind,
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Fatal(err)
	}
}
