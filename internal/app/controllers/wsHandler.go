package controllers

import (
	"fmt"
	"forum/internal/app/websok"
	"net/http"

	"github.com/gorilla/websocket"
)

// same struct websocket upgrader + handleconnec + gochatconnection

// websocket upgrader
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}
// var ClientsMap = make(map[string]*websok.Client)
// hanle websocketconnection
func HandleConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("err upgrading websocket ", err)
		return
	}

	/// i shold extract username from query

	// crieate new client

	// it shild be the username so fetch aftes session and extract username from query
	usernname := r.URL.Query().Get("username")

	client := &websok.Client{
		Conn:     conn,
		Username: usernname, // Extract username from query
		Send:     make(chan websok.Privatemessagestruct),
	}

	///regester client use this for ping pong

	websok.ChatHub.Regester <- client

	// ChatHub.regester<-Client

	// start reading and writing message  go

	go client.ReadMessages()
	go client.Writemessages()
}
