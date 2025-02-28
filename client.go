package main

import (
	// "runtime/metrics"

	"github.com/gorilla/websocket"
)

//clien represent a single chating user.

type client struct{
	//socket is the web socket for this client 
	socket *websocket.Conn

	//recive is a channel to resive message from the athers users 

	receive chan []byte
	
	//rome is the rome that the clien is messaging in 

	room *room

}









func (c *client)read(){

	defer c.socket.Close()

	for{

		// _,msg,err:=c.socket.ReadMessage()
		
	_,msg,err:=c.socket.ReadMessage()
		if err!=nil{
			return
			
		}
		c.room.forward<-msg

	}
}



func (c *client)write() {
	defer c.socket.Close()
	
	for msg:=range c.receive{
		err:=c.socket.WriteMessage(websocket.TextMessage,msg)

		if err!=nil{
			return
		}

	
	}
	
}