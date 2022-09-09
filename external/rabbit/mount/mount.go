package mountRabbit

import (
	"golangNetHttp/external/rabbit/bind"
	rabbitConfig "golangNetHttp/external/rabbit/config"
	"golangNetHttp/external/rabbit/exchange"
	"golangNetHttp/external/rabbit/queues"
)

func Mount() {
	ch := rabbitConfig.CreateChannel()

	exchange.Declare(ch, string(exchange.Exchanges["GREEN"]), string(exchange.KindExchanges["FANOUT"]))
	q := queues.Declare(ch, string(queues.Queues["GREEN-QUEUE"]))
	bind.BindQueueRoutingExchange(ch, q.Name, string(bind.Routingkies["GREEN-ROUTING"]), string(exchange.Exchanges["GREEN"]))
	rabbitConfig.SendMsgRabbit(ch, string(exchange.Exchanges["GREEN"]), "Hello rabbit 6", string(bind.Routingkies["GREEN-ROUTING"]))
}
