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
	Clients        map[string]*Client
	Regester       chan *Client
	Unregester     chan *Client
	Messages       chan Privatemessagestruct
	LastInertedMsg int64
	mu             sync.Mutex
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
		select {
		case client := <-h.Regester:
			fmt.Println(<-h.Regester, "ddddd")
			fmt.Println("!!!!!!!!!!!!!!!!!")
			h.mu.Lock()
			h.Clients[client.Username] = client
			h.mu.Unlock()

		case client := <-h.Unregester:
			fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!1", client.Username)
			h.mu.Lock()
			delete(h.Clients, client.Username)
			h.mu.Unlock()
			client.Conn.Close()

		case msg := <-h.Messages:

			h.mu.Lock()
			// show cliant
			receiver, exists := h.Clients[msg.Receiver]
			h.mu.Unlock()

			if exists {

				L, err := SaveMessage(utils.Db1.Db, msg.Sender, msg.Receiver, msg.Content)
				if err != nil {
					fmt.Println("err save mesage", err)
				}
				h.LastInertedMsg = L
				//fmt.Println("---------------", O)
				receiver.Send <- msg

			} else {
				fmt.Println("Receiver not found:", msg.Receiver)
			}
		}
	}
}

func (h *Storeactivewebsocketclient) GetOnlineUsersnames() map[string]bool {
	var online = make(map[string]bool)

	for username := range h.Clients {
		online[username] = true
	}
	return online
}

func UnregisterClientByUsername(username string) bool {
	ChatHub.mu.Lock()
	defer ChatHub.mu.Unlock()

	client, exists := ChatHub.Clients[username]
	if exists {
		ChatHub.Unregester <- client
		return true
	} else {
		for v := range ChatHub.Clients {
			fmt.Println("stay", v)
		}
	}
	return false
}
