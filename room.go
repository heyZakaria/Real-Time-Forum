package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type room struct {
	// clients holds all current client in this room

	clients map[*client]bool

	// joint is a channel for client wishing to join the room.

	joint chan *client

	//  left is a channel for client wishing to left the room.

	leave chan *client

	// forward is a channels that holds incomming messages that should be forwarded to the other client .

	forward chan []byte
}

// func newroom() *room {
// 	return &room{
// 		forward: make(chan []byte),
// 		joint:   make(chan *client),
// 		leave:   make(chan *client),
// 		clients: make(map[*client]bool),
// 	}
// }

func newRoom() *room {
	return &room{
		forward: make(chan []byte),
		joint:   make(chan *client),
		leave:   make(chan *client),
		clients: make(map[*client]bool),
	}
}

func (r *room) run() {
	for {
		select {
		case client := <-r.joint:
			r.clients[client] = true

		case client := <-r.leave:
			// r.client[client]=false
			delete(r.clients, client)

			close(client.receive)

		case msg := <-r.forward:

			for clients := range r.clients {
				clients.receive <- msg
			}

		}
	}
}

const (
	socketBufferSize  = 1024
	messageBufferSize = 1024
)

var upgrader = &websocket.Upgrader{ReadBufferSize: socketBufferSize, WriteBufferSize: socketBufferSize}

func (r *room) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal("servehttp err", err)

		return
	}

	client := &client{
		socket:  socket,
		receive: make(chan []byte, messageBufferSize),
		room:    r,
	}

	r.joint <- client
	defer func() { r.leave <- client }()

	go client.write()

	client.read()
}
