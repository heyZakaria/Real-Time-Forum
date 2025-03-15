package websok

import (
	"fmt"
	"sync"

	"forum/internal/app/models/utils"
)

// edea is struct store the websocket (client:keyusrnam) >> shold update on struct like the hup in last commit ...
// i have hup in package websok
// but GetOnlineUsersHandler() in package controller
type Storeactivewebsocketclient struct {
	Clients    map[string]*Client
	Regester   chan *Client
	Unregester chan *Client
	Messages   chan Privatemessagestruct
	mu         sync.Mutex
}

var ChatHub = &Storeactivewebsocketclient{
	Clients:    make(map[string]*Client),
	Regester:   make(chan *Client),
	Unregester: make(chan *Client),
	Messages:   make(chan Privatemessagestruct),
}

type Privatemessagestruct struct {
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
	Content  string `json:"content"`
}

/// run  listen for regestered private message

func (h *Storeactivewebsocketclient) Run() {
	for {
		// fmt.Println(h.Clients, "all my clients")
		// /	Clients:    make(map[string]*Client),

		// for username := range h.Clients {
		// 	fmt.Println("add to on_line user ", username)
		// }

		select {
		case client := <-h.Regester:
			h.mu.Lock()
			h.Clients[client.Username] = client
			h.mu.Unlock()
			fmt.Println(client.Username, "was connected.")

		case client := <-h.Unregester:
			h.mu.Lock()
			delete(h.Clients, client.Username)
			h.mu.Unlock()
			client.Conn.Close()

			///need to add ping pong logic and add it to mt handler
			fmt.Println(client.Username, " by by close xanel disconnected")

		case msg := <-h.Messages:
			fmt.Println("New message received in ChatHub:", msg)

			///////////////////////////////////////////////////////////lets save message to db

			val, err := SaveMessage(utils.Db1.Db, msg.Sender, msg.Receiver, msg.Content)
			if err != nil {
				fmt.Println("err save mesage", err)
			}
			fmt.Println("vaaal last indicat", val)

			h.mu.Lock()
			// show cliant
			fmt.Println("Connected clients:", h.Clients)
			receiver, exists := h.Clients[msg.Receiver]
			h.mu.Unlock()

			if exists {
				fmt.Println("snd to:", msg.Receiver)
				receiver.Send <- msg
			} else {
				fmt.Println("Receiver not found:", msg.Receiver)
			}

		}
	}
}

func (h *Storeactivewebsocketclient) GetOnlineUsersnames() []string {
	h.mu.Lock()
	defer h.mu.Unlock()
	usernames := make([]string, 0, len(h.Clients))
	for username := range h.Clients {
		usernames = append(usernames, username)
	}
	return usernames
}
