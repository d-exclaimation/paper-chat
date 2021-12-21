//
//  event_emitter.go
//  streaming
//
//  Created by d-exclaimation on 11:10 PM.
//  Copyright © 2021 d-exclaimation. All rights reserved.
//

package emitter

import (
	"context"
	"github.com/d-exclaimation/paper-chat/graphql/model"
)

type Void = struct{}

type EventEmitter struct {
	// -- Private states --
	consumers map[chan *model.Message]Void

	// -- Channels --
	onRegister     chan chan *model.Message
	onTerminate chan chan *model.Message
	onMessage chan *model.Message
	onAcid    chan Void
}

func New() *EventEmitter {
	emitter := &EventEmitter{
		consumers:   make(map[chan *model.Message]Void),
		onRegister:  make(chan chan *model.Message),
		onTerminate: make(chan chan *model.Message),
		onMessage:   make(chan *model.Message),
		onAcid:      make(chan Void),
	}
	go emitter.onReceive()
	return emitter
}

func (e *EventEmitter) onReceive() {
	for {
		select {
		case registeredConsumer := <-e.onRegister:
			e.consumers[registeredConsumer] = Void{}
		case terminatedConsumer := <-e.onTerminate:
			delete(e.consumers, terminatedConsumer)
		case message := <-e.onMessage:
			for consumer, _ := range e.consumers {
				consumer <- message
			}
		case <-e.onAcid:
			for consumer, _ := range e.consumers {
				close(consumer)
				delete(e.consumers, consumer)
			}
		}
	}
}

func (e *EventEmitter) Consumer(ctx context.Context) chan *model.Message {
	consumer := make(chan *model.Message)
	e.onRegister <- consumer

	go func() {
		<-ctx.Done()
		e.onTerminate <- consumer
	}()

	return consumer
}

func (e *EventEmitter) Produce(message *model.Message) {
	e.onMessage <- message
}

func (e *EventEmitter) End() {
	e.onAcid <- Void{}
}