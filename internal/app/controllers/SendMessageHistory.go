package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"forum/internal/app/models/utils"
	"forum/internal/app/websok"
)

type Message struct {
	ID         int64
	SenderID   string
	ReceiverID string
	Content    string
	CreatedAt  time.Time
}

func SendMessageHistory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	user1 := r.URL.Query().Get("user1")
	user2 := r.URL.Query().Get("user2")

	if user1 == "" || user2 == "" {
		/// Bad request
		fmt.Println("no user found")
		return
	}

	conversation, err := websok.GetConversationHistory(utils.Db1.Db, user1, user2, 100)
	if err != nil {
		fmt.Println("err getting convertation")
		return
	}
	err = json.NewEncoder(w).Encode(conversation)
	if err != nil {
		fmt.Println("cant encode response convertation")
		return
	}
}
