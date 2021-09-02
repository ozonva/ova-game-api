package kafka

// MessageType - represents type for iota with Create, MultiCreate, Update, Delete operations
type MessageType int

const (
	// Create - Create  via Producer
	Create MessageType = 1
	// MultiCreate - Create several s via Producer
	MultiCreate MessageType = 2
	// Update - Update  via Producer
	Update MessageType = 3
	// Delete - Delete  via Producer
	Delete MessageType = 4
	// Describe - Describe  via Producer
	Describe MessageType = 4
)

// Message - message for Kafka
type Message struct {
	MessageType MessageType
	Value       interface{}
}
