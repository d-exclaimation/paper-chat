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
	registerChannel  chan chan *model.Message
	terminateChannel chan chan *model.Message
	messageChannel   chan *model.Message
	acidChannel      chan Void
}

func New() *EventEmitter {
	emitter := &EventEmitter{
		consumers:        make(map[chan *model.Message]Void),
		registerChannel:  make(chan chan *model.Message),
		terminateChannel: make(chan chan *model.Message),
		messageChannel:   make(chan *model.Message),
		acidChannel:      make(chan Void),
	}
	go emitter.onMessage()
	return emitter
}

func (e *EventEmitter) onMessage() {
	for {
		select {
		case registeredConsumer := <-e.registerChannel:
			e.consumers[registeredConsumer] = Void{}
		case terminatedConsumer := <-e.terminateChannel:
			delete(e.consumers, terminatedConsumer)
		case message := <-e.messageChannel:
			for consumer, _ := range e.consumers {
				consumer <- message
			}
		case <-e.acidChannel:
			for consumer, _ := range e.consumers {
				close(consumer)
				delete(e.consumers, consumer)
			}
		}
	}
}

func (e *EventEmitter) Consumer(ctx context.Context) chan *model.Message {
	consumer := make(chan *model.Message)
	e.registerChannel <- consumer

	go func() {
		<-ctx.Done()
		e.terminateChannel <- consumer
	}()

	return consumer
}

func (e *EventEmitter) Produce(message *model.Message) {
	e.messageChannel <- message
}

func (e *EventEmitter) End() {
	e.acidChannel <- Void{}
}