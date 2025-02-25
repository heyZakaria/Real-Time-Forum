package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var websocketupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Manager struct {
	clients ClientList
	sync.RWMutex

	handlers map[string]EventHandler
}

func newmanager() *Manager {
	m := &Manager{
		clients:  make(ClientList),
		handlers: make(map[string]EventHandler),
	}
	m.setupEventHandlers()
	return m
}

func (m *Manager) setupEventHandlers() {
	m.handlers[EventSendMessage] = SendMessage
}


func SendMessage(event Event, c *Client) error {
	fmt.Println(event)
	return nil
}



func (m *Manager) routeEvent(event Event, c *Client) error {
	//chek if event type 
	if handler, ok := m.handlers[event.Type]; ok {
		if err := handler(event, c); err != nil {
			return err
		}
		return nil
	} else {
		return errors.New("thie is no such event type")
	}
}

func (m *Manager) serveWS(w http.ResponseWriter, r *http.Request) {
	conn, err := websocketupgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	// conn.Close()
	client := NewClient(conn, m)
	m.addclient(client)

	// startt

	go client.readMessages()
	go client.WriteMessage()
}

func (m *Manager) addclient(client *Client) {
	m.Lock()
	defer m.Unlock()

	// if _,ok:=m.clients[client];ok{

	// }

	m.clients[client] = true
}

func (m *Manager) removclient(client *Client) {
	m.Lock()
	defer m.Unlock()

	if _, ok := m.clients[client]; ok {
		client.connection.Close()
		delete(m.clients, client)
	}

	// m.clients[client]=false
}
