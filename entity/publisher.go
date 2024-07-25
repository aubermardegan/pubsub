package entity

import (
	"context"
	"fmt"
	"sync"

	"github.com/google/uuid"
)

var (
	ErrTopicNotFound = fmt.Errorf("topic not found")
)

type Topic string

type Publisher struct {
	sync.RWMutex
	subs         map[Topic]map[SubId]*Subscription
	doneListener chan SubId
}

func NewPublisher() *Publisher {
	return &Publisher{
		subs:         make(map[Topic]map[SubId]*Subscription),
		doneListener: make(chan SubId),
	}
}

func (p *Publisher) Start(ctx context.Context) {
	fmt.Printf("\nStarting Publisher")
	for {
		select {
		case subId := <-p.doneListener:
			p.Lock()
			for _, subsInTopic := range p.subs {
				sub, ok := subsInTopic[subId]
				if ok {
					sub.quit <- struct{}{}
					close(sub.messages)
					delete(subsInTopic, subId)
				}
			}
			fmt.Printf("\nSubscriber Id %s removed", uuid.UUID(subId))
			p.Unlock()
		case <-ctx.Done():
			fmt.Printf("\nClosing Publisher")
			return
		}
	}
}

func (p *Publisher) Publish(t Topic, message string) error {
	p.RLock()
	defer p.RUnlock()

	subs, ok := p.subs[t]
	if !ok {
		return ErrTopicNotFound
	}

	var wg sync.WaitGroup
	for _, sub := range subs {
		wg.Add(1)
		go func(*sync.WaitGroup) {
			defer wg.Done()
			sub.messages <- message
		}(&wg)
	}
	wg.Wait()
	return nil
}

func (p *Publisher) Subscribe(t Topic) *Subscription {
	sub := NewSubscription(p.doneListener)
	p.Lock()
	defer p.Unlock()

	topicMap, ok := p.subs[t]
	if !ok {
		topicMap = make(map[SubId]*Subscription)
	}
	topicMap[sub.Id] = sub
	p.subs[t] = topicMap
	return sub
}
