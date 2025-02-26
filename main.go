package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}


func treader(conn *websocket.Conn){
	for {
		messageType,p,err:=conn.ReadMessage()
		if err!=nil{
			log.Println(err)
			return 
			

		}
	log.Println(string(p))

	if err:=conn.WriteMessage(messageType,p);err!=nil{
		log.Println(err)
		return
	}

	}


}


func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "home page ")


}

func wsEndPoint(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprint(w, "wsendpoint")
	upgrader.CheckOrigin=func(r *http.Request) bool {return true }

	ws,err:=upgrader.Upgrade(w,r,nil)
	if err!=nil{
		log.Println(err,"Errwsendpoint")
	}
	log.Println("CLien succefully...")
	treader(ws)
}

func setupRoutes() {
	http.HandleFunc("/", homepage)
	http.HandleFunc("/ws", wsEndPoint)
}

func main() {
	fmt.Println("helo wbsocket")
	setupRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
