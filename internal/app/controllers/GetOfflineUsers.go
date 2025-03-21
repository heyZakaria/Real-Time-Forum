package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"maps"
	"net/http"

	"forum/internal/app/models/utils"
	"forum/internal/app/websok"
)

func GetOfflineUsersHandler(w http.ResponseWriter, r *http.Request) {
	onlinUsernames := websok.ChatHub.GetOnlineUsersnames()
	allUsers, err := allUsers(utils.Db1.Db)
	if err != nil {
		http.Error(w, "username not fund", http.StatusInternalServerError)
	}

	token, err := r.Cookie("session_id")
	if err != nil {
		return
	}
	user_id := 0
	if token != nil {

		user_id, err = SelectUser(token.Value)
		if err != nil {
			http.Error(w, "username not fund", http.StatusInternalServerError)
			
		}
	}
	curentuser, _ := SelectUsername(user_id)
	
	
	lastTalked, err := allUsersByLastMSG(utils.Db1.Db, curentuser)
	if err != nil {
		http.Error(w, "username not fund", http.StatusInternalServerError)
	}

	offlineUsers := offlinePeople(onlinUsernames, allUsers)

	// This copies all onlineusername to offlineUsers
	// so we have one object, I need sorting
	maps.Copy(offlineUsers, onlinUsernames)

	// this is just one object with 2 objects inside it
	// you can access them in frontend (ws.js in fetchOfflineUsers function)

	finalMap := map[string]any{
		"allUsers":   offlineUsers,
		"lastTalked": lastTalked,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(finalMap)
}

func allUsersByLastMSG(db *sql.DB, curentuser string) (map[string]int, error) {
	query := `SELECT sender_id, receiver_id, created_at FROM messages ORDER BY created_at DESC`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	lastTalked := make(map[string]int)
	var sender, receiver, created_at string
	count := 0

	for rows.Next() {
		if err := rows.Scan(&sender, &receiver, &created_at); err != nil {
			fmt.Println("err scanning messages:", err)
			continue
		}

		if curentuser == sender || curentuser == receiver {
			if _, ok := lastTalked[sender]; !ok {
				lastTalked[sender] = count
				count++
			}

			if _, ok := lastTalked[receiver]; !ok {
				lastTalked[receiver] = count
				count++
			}
		}
	}

	return lastTalked, nil
}

func allUsers(db *sql.DB) ([]string, error) {
	query := `SELECT username FROM users ORDER BY username`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []string
	for rows.Next() {
		var user string

		if err := rows.Scan(&user); err != nil {
			fmt.Println("err scaning users ,user", err)
		}
		users = append(users, user)
	}
	return users, nil
}

func offlinePeople(online map[string]bool, allUsers []string) map[string]bool {
	offline := make(map[string]bool)
	for _, username := range allUsers {
		if notContains(online, username) {
			offline[username] = false
		}
	}
	return offline
}

func notContains(online map[string]bool, username string) bool {
	for v := range online {
		if v == username {
			return false
		}
	}
	return true
}
