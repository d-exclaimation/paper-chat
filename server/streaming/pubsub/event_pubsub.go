//
//  event_pubsub.go
//  pubsub
//
//  Created by d-exclaimation on 11:28 PM.
//  Copyright © 2021 d-exclaimation. All rights reserved.
//

package pubsub

import (
	"context"
	"github.com/d-exclaimation/paper-chat/graphql/model"
	"github.com/d-exclaimation/paper-chat/streaming/emitter"
)

type streamRequest struct {
	trigger string
	channel chan *emitter.EventEmitter
}

type publish struct {
	trigger string
	message *model.Message
}

type EventPubSub struct {
	// -- Private states --
	emitters map[string]*emitter.EventEmitter

	// -- Channels
	onStream chan streamRequest
	onPublish chan publish
	onClose chan string
}

func New() *EventPubSub {
	pubsub := &EventPubSub{
		emitters:  make(map[string]*emitter.EventEmitter),
		onStream:  make(chan streamRequest),
		onPublish: make(chan publish),
		onClose:   make(chan string),
	}
	go pubsub.onReceive()
	return pubsub
}

func (e *EventPubSub) onReceive() {
	for {
		select {
		case newConsumer := <-e.onStream:
			emit, ok := e.emitters[newConsumer.trigger]
			if !ok {
				emit = emitter.New()
			}

			newConsumer.channel <- emit

			e.emitters[newConsumer.trigger] = emit
		case published :=  <-e.onPublish:
			emit, ok := e.emitters[published.trigger]
			if ok {
				emit.Produce(published.message)
			}
		case trigger := <-e.onClose:
			emit, ok := e.emitters[trigger]
			if ok {
				emit.End()
			}
			delete(e.emitters, trigger)
		}
	}
}

func (e *EventPubSub) Stream(trigger string, ctx context.Context) chan *model.Message {
	futureEmitter := make(chan *emitter.EventEmitter)
	e.onStream <- streamRequest{
		trigger: trigger,
		channel: futureEmitter,
	}
	emit := <-futureEmitter
	return emit.Consumer(ctx)
}

func (e *EventPubSub) Publish(trigger string, message *model.Message) {
	e.onPublish <- publish{
		trigger: trigger,
		message: message,
	}
}

func (e *EventPubSub) Close(trigger string) {
	e.onClose <- trigger
}