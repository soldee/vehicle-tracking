package services

import (
	"fmt"
	"sync"
)

type Sub struct {
	id     int
	active bool
	msgs   chan string
	errors chan error
	mutex  sync.RWMutex
}

type Broker struct {
	subs      []*Sub
	freeSlots chan int
	mutex     sync.RWMutex
}

func NewBroker() *Broker {
	return &Broker{
		subs:      make([]*Sub, 0),
		freeSlots: make(chan int, 1000),
	}
}

func (b *Broker) Subscribe() *Sub {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	var sub *Sub
	if len(b.freeSlots) != 0 {
		sub = b.subs[<-b.freeSlots]
		sub.mutex.Lock()
		sub.active = true
		sub.mutex.Unlock()
	} else {
		sub = &Sub{
			id:     len(b.subs),
			active: true,
			msgs:   make(chan string),
			errors: make(chan error),
		}
		b.subs = append(b.subs, sub)
	}

	return sub
}

func (b *Broker) Unsubscribe(id int) error {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	if id > len(b.subs) {
		return fmt.Errorf("id %v is not subscribed", id)
	}

	sub := b.subs[id]

	sub.mutex.Lock()
	defer sub.mutex.Unlock()

	sub.active = false
	b.freeSlots <- id
	return nil
}

func (b *Broker) Publish(msg string) {
	for _, sub := range b.subs {
		sub.mutex.Lock()
		if sub.active {
			sub.msgs <- msg
		}
		sub.mutex.Unlock()
	}
}

func (b *Broker) PublishError(err error) {
	for _, sub := range b.subs {
		sub.mutex.Lock()
		if sub.active {
			sub.errors <- err
		}
		sub.mutex.Unlock()
	}
}
