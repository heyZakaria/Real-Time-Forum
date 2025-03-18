package websok

import (
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
)

// Client by definition is just a websocket connection
type Client struct {
	Conn     *websocket.Conn
	Username string
	Send     chan Privatemessagestruct
}

// reade message its just listning for incoming messages
func (c *Client) ReadMessages() {
	defer func() {
		ChatHub.Unregester <- c
		c.Conn.Close()
	}()

	for {

		var msg Privatemessagestruct
		err := c.Conn.ReadJSON(&msg)
		if err != nil {
			fmt.Println("err erading msg mfk", err)
			break
		}
		// solve the problem of the sender dammmnit
		msg.Sender = c.Username
		ChatHub.Messages <- msg

	}
}

// writemessage send message to the client
func (c *Client) Writemessages() {
	for msg := range c.Send {

		msgjson, _ := json.Marshal(msg)
		fmt.Println(string(msgjson), "waaaaaaaa")
		err := c.Conn.WriteMessage(websocket.TextMessage, msgjson)
		if err != nil {
			fmt.Println("err writng the message mfk", err)
			break

		}

	}
}
