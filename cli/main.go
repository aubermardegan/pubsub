package main

import (
	"context"
	"time"

	"github.com/aubermardegan/pubsub/entity"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	publisher := entity.NewPublisher()
	go publisher.Start(ctx)

	sub1 := publisher.Subscribe("animals")
	sub2 := publisher.Subscribe("animals")
	sub3 := publisher.Subscribe("fish")

	go sub1.Listen()
	go sub2.Listen()
	go sub3.Listen()

	publisher.Publish("animals", "meow")
	publisher.Publish("animals", "woof")
	sub1.Unsubscribe()
	sub2.Unsubscribe()

	time.Sleep(2 * time.Second)          //tempo para fechar a subscrição dos canais
	publisher.Publish("animals", "roar") //nenhuma subscrição para receber essa mensagem
	publisher.Publish("fish", "splash")

	time.Sleep(2 * time.Second)
	cancel()
	time.Sleep(2 * time.Second)
}
