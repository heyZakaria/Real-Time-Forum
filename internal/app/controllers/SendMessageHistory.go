package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
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

	limit := 10
	if strlimit := r.URL.Query().Get("limit"); strlimit != "" {
		ParsedLimit, err := strconv.Atoi(strlimit)
		if err == nil && ParsedLimit > 0 {
			limit = ParsedLimit
		} else {
			fmt.Println(err)
		}
	}

	// Parse offset parameter with default value of 0
	offset := 0
	if stroffset := r.URL.Query().Get("offset"); stroffset != "" {
		ParsedOffset, err := strconv.Atoi(stroffset)
		if err == nil && ParsedOffset >= 0 {
			offset = ParsedOffset
		} else {
			fmt.Println(err)
		}
	}

	conversation, err := websok.GetConversationHistory(utils.Db1.Db, user1, user2, limit, offset)
	if err != nil {
		http.Error(w, "Failed to get conversation", http.StatusInternalServerError)

		return
	}
	//fmt.Println(conversation[len(conversation)-1], "zzzzzzz")
	err = json.NewEncoder(w).Encode(conversation)
	if err != nil {
		http.Error(w, "cant encode response convertation", http.StatusInternalServerError)

		return
	}
}
