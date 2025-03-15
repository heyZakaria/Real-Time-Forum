package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"forum/internal/app/models/utils"
	"forum/internal/app/websok"
)

func GetOfflineUsersHandler(w http.ResponseWriter, r *http.Request) {
	var offlineUsers []string
	onlinusernames := websok.ChatHub.GetOnlineUsersnames()
	allUsers, err := allUsers(utils.Db1.Db)
	if err != nil {
		fmt.Println(err)
	}

	offlineUsers = offlinepeaple(onlinusernames, allUsers)

	type UserResponse struct {
		Username string `json:"username"`
	}
	response := make([]UserResponse, len(offlineUsers))

	for i, username := range offlineUsers {
		response[i] = UserResponse{
			Username: username,
		}
	}

	//fmt.Println("XXX", offlineUsers)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
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

func offlinepeaple(online []string, allUsers []string) []string {
	var offline []string
	for _, y := range allUsers {
		if notContains(online, y) {
			offline = append(offline, y)
		}
	}

	return offline
}

func notContains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return false
		}
	}
	return true
}
