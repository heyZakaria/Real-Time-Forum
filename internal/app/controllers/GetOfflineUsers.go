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

	onlinusernames := websok.ChatHub.GetOnlineUsersnames()
	allUsers, err := allUsers(utils.Db1.Db)
	if err != nil {
		fmt.Println(err)
	}

	offlineUsers := offlinepeaple(onlinusernames, allUsers)

	
	// This copies all onlineusername to offlineUsers
	// so we have one object, I need sorting
	maps.Copy(offlineUsers, onlinusernames)
	fmt.Println(offlineUsers)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(offlineUsers)
}

func allUsers(db *sql.DB) ([]string, error) {
	query := `SELECT username FROM users`
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

func offlinepeaple(online map[string]bool, allUsers []string) map[string]bool {
	var offline = make(map[string]bool)
	for _, username := range allUsers {
		if notContains(online, username) {
			offline[username] = false
		}
	}

	return offline
}

func notContains(s map[string]bool, str string) bool {
	for v := range s {
		if v == str {
			return false
		}
	}
	return true
}
