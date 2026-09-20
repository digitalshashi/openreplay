package handlers

import . "openreplay/backend/pkg/messages"

type MessageProcessor interface {
	Handle(message Message, timestamp uint64) Message
	Build() Message
	MessageTypes() []int
}
