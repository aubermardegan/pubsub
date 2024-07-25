package entity

import (
	"fmt"

	"github.com/google/uuid"
)

type SubId uuid.UUID

type Subscription struct {
	Id       SubId
	messages chan string
	done     chan SubId
	quit     chan struct{}
}

func NewSubscription(doneListener chan SubId) *Subscription {
	return &Subscription{
		Id:       SubId(uuid.New()),
		messages: make(chan string),
		done:     doneListener,
		quit:     make(chan struct{}),
	}
}

func (s *Subscription) Listen() {
	for {
		select {
		case msg := <-s.messages:
			fmt.Printf("\nMensagem Recebida no Sub %s: %s", uuid.UUID(s.Id), msg)
		case <-s.quit:
			return
		}
	}
}

func (s *Subscription) Unsubscribe() {
	s.done <- s.Id
}
