package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"forum/internal/app/controllers"
	"forum/internal/app/models/api"
	"forum/internal/app/models/utils"
	"forum/internal/app/websok"
	"forum/internal/platform/database"

	"github.com/gorilla/websocket"
	_ "github.com/mattn/go-sqlite3"
)

// same struct websocket upgrader + handleconnec + gochatconnection

// websocket upgrader
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

// hanle websocketconnection
func handleConnection(w http.ResponseWriter, r *http.Request) {
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

func main() {
	db, _ := database.CreateDatabase()

	go websok.ChatHub.Run()

	http.HandleFunc("/ws", handleConnection)

	utils.Db1.Db = db

	http.HandleFunc("/internal/app/views/static/", SetupStaticFilesHandler)

	http.HandleFunc("/", controllers.Home)

	/* http.HandleFunc("/register", controllers.Registration)
	http.HandleFunc("/login", controllers.Login)
	*/
	http.HandleFunc("/api", api.Codage)
	http.HandleFunc("/api/registred", controllers.CheckRegistration)
	http.HandleFunc("/api/online-users", controllers.GetOnlineUsersHandler)
	http.HandleFunc("/api/current-user",controllers.GetCurrentUsername)

	http.HandleFunc("/api/addComment", controllers.AddComment)
	http.HandleFunc("/api/addPost", controllers.AddPost)

	http.HandleFunc("/api/posts/", controllers.HandleReaction)
	http.HandleFunc("/api/comments/", controllers.HandleCommentReaction)

	http.HandleFunc("/api/logout", controllers.Logout)

	fmt.Println("server runing at http://localhost:4444")
	err := http.ListenAndServe(":4444", nil)
	if err != nil {
		fmt.Println(err, "we can't serve")
		return
	}
}

func SetupStaticFilesHandler(w http.ResponseWriter, r *http.Request) {
	staticDir := "./internal/app/views/static"

	afterStatic := r.URL.Path[len("/internal/app/views/static/"):]
	if strings.HasSuffix(afterStatic, "/") {
		utils.MessageError(w, r, http.StatusNotFound, "What are you doing here!")
		return
	}

	fullPath := filepath.Join(staticDir, afterStatic)

	fileinfo, err := os.Stat(fullPath) // get the endpoint file infos: name size adress...
	if err != nil || fileinfo.IsDir() {
		utils.MessageError(w, r, http.StatusNotFound, "What are you doing here!")
		return
	}

	http.ServeFile(w, r, fullPath)
}
