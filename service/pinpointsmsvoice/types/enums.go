// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type EventType string

// Enum values for EventType
const (
	EventTypeInitiatedCall EventType = "INITIATED_CALL"
	EventTypeRinging       EventType = "RINGING"
	EventTypeAnswered      EventType = "ANSWERED"
	EventTypeCompletedCall EventType = "COMPLETED_CALL"
	EventTypeBusy          EventType = "BUSY"
	EventTypeFailed        EventType = "FAILED"
	EventTypeNoAnswer      EventType = "NO_ANSWER"
)

// Values returns all known values for EventType. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (EventType) Values() []EventType {
	return []EventType{
		"INITIATED_CALL",
		"RINGING",
		"ANSWERED",
		"COMPLETED_CALL",
		"BUSY",
		"FAILED",
		"NO_ANSWER",
	}
}
