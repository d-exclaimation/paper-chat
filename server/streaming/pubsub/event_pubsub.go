//
//  event_pubsub.go
//  pubsub
//
//  Created by d-exclaimation on 11:28 PM.
//  Copyright © 2021 d-exclaimation. All rights reserved.
//

package pubsub

import "github.com/d-exclaimation/paper-chat/streaming/emitter"

type EventPubSub struct {
	// -- Private states --
	emitters map[*emitter.EventEmitter]emitter.Void
}
