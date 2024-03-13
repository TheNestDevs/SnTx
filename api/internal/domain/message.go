package domain

// Message holds the structure of JSON message send via websocket.
type Message struct {
	ClientID string `json:"client_id,omitempty"` // client id of the client
	Topic    string `json:"topic,omitempty"`     // topic of the message sent
	Msg      string `json:"message,omitempty"`   // message string that is sent
}
