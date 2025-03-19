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

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, _ := database.CreateDatabase()

	go websok.ChatHub.Run()

	http.HandleFunc("/ws", controllers.HandleConnection)

	utils.Db1.Db = db

	http.HandleFunc("/internal/app/views/static/", SetupStaticFilesHandler)

	http.HandleFunc("/", controllers.Home)

	/* http.HandleFunc("/register", controllers.Registration)
	http.HandleFunc("/login", controllers.Login)
	*/
	http.HandleFunc("/api", api.Codage)
	http.HandleFunc("/api/registred", controllers.CheckRegistration)

	// http.HandleFunc("/api/online-users", controllers.GetOnlineUsersHandler)
	http.HandleFunc("/api/friends-list", controllers.GetOfflineUsersHandler)
	http.HandleFunc("/api/message-history", controllers.SendMessageHistory)
	http.HandleFunc("/api/sendnotification",controllers.Sendnotification)
	// http.HandleFunc("/api/sendnotification",controllers.Sendnotification)
	http.HandleFunc("/api/current-user", controllers.GetCurrentUsername)

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
