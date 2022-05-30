package request

// Publish is a request to publish a message to a topic.
type Publish struct {
	Topic   string `json:"topic,omitempty"`
	Message string `json:"message,omitempty"`
}
