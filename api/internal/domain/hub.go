package domain

type Room struct {
	Topic   string
	Clients map[string]*Client
}

type Hub struct {
	Rooms     map[string]*Room // Map of rooms. Key is the Topic and value is the Room
	Broadcast chan *Message
	NewClient chan *Client
}

func NewHub() *Hub {
	return &Hub{
		Rooms:     make(map[string]*Room),
		Broadcast: make(chan *Message),
		NewClient: make(chan *Client),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case message := <-h.Broadcast:
			h.SendToRoom(message.Topic, message)
		case newClient := <-h.NewClient:
			if _, ok := h.Rooms[newClient.Topic]; !ok {
				h.Rooms[newClient.Topic] = &Room{
					Topic:   newClient.Topic,
					Clients: make(map[string]*Client),
				}
				h.Rooms[newClient.Topic].Clients[newClient.Id] = newClient
			} else {
				h.Rooms[newClient.Topic].Clients[newClient.Id] = newClient
			}
		}
	}
}

// SendToRoom sends a message to all clients in a room accept the sender client
func (h *Hub) SendToRoom(room string, message *Message) {
	for _, client := range h.Rooms[room].Clients {
		if client.Id != message.ClientID {
			client.Conn.WriteJSON(message)
		}
	}
}

func (h *Hub) RemoveClient(client *Client) {
	delete(h.Rooms[client.Topic].Clients, client.Id)

	if len(h.Rooms[client.Topic].Clients) == 0 {
		delete(h.Rooms, client.Topic)
	}
}
