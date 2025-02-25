package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

var (
	pongwait = 10 * time.Second

	pinginterval = (pongwait * 9) / 10
)



type ClientList map[*Client]bool

type Client struct {
	connection *websocket.Conn
	manager    *Manager
	// egress used to avoid concurrent writes on the websocket connection

	egress chan Event
}

func NewClient(conn *websocket.Conn, manager *Manager) *Client {
	return &Client{
		connection: conn,
		manager:    manager,
		egress:     make(chan Event),
	}
}

func (c *Client) readMessages() {
	defer func() {
		// clean up conn

		c.manager.removclient(c)
	}()

	for {
		_, payload, err := c.connection.ReadMessage()
		if err != nil {
			// break
			if websocket.IsUnexpectedCloseError(err, websocket.CloseAbnormalClosure, websocket.CloseGoingAway) {
				log.Println("err redding messages:%", err)
			}
			break
		}
		////brodcasr to all the peaple
		// 	for wsclient:=range c.manager.clients{
		// 		wsclient.egress<-payload
		// 	}

		// log.Println(messagetype)
		// log.Println(string(payload))

		var request Event

		if err := json.Unmarshal(payload, &request); err != nil {
			log.Println("err marshaling ebvent %v", err)
			break
		}

		if err := c.manager.routeEvent(request, c); err != nil {
			log.Println("err handeling message", err)
		}

	}
}

func (c *Client) WriteMessage() {
	defer func() {
		c.manager.removclient(c)
	}()


	tiker:=time.NewTicker(pinginterval)
	for {
		select {
		case messages, ok := <-c.egress:
			if !ok {
				if err := c.connection.WriteMessage(websocket.CloseMessage, nil); err != nil {
					log.Println("connection closed", err)
				}
				return
			}

			data, err := json.Marshal(messages)
			if err != nil {
				log.Println(err)
				return
			}

			if err := c.connection.WriteMessage(websocket.TextMessage, data); err != nil {
				log.Println("failed to send mesages", err)
			}
			log.Println("messages send")
		case <-tiker.C:
			log.Println("piNG")
			//send ping to te client

			if  err :=c.connection.WriteMessage(websocket.PingMessage,[]byte(``));err!=nil{

				log.Println("wetemesg err",err)
				return
			}
		
		}

	}
}
