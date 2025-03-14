package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"forum/internal/app/models/utils"
	"forum/internal/app/websok"
)

func GetOflineUsersHandler(w http.ResponseWriter, r *http.Request) {
	var oflineusers []string
	onlinusernames := websok.ChatHub.GetOnlineUsersnames()
	allusers, err := allusers(utils.Db1.Db)
	if err != nil {
		fmt.Println(err)
	}

	oflineusers = offlinepeaple(onlinusernames, allusers)

	type UserResponse struct {
		Username string `json:"username"`
	}
	response := make([]UserResponse, len(oflineusers))

	for i, username := range oflineusers {
		response[i] = UserResponse{
			Username: username,
		}
	}

	fmt.Println("XXX", oflineusers)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func allusers(db *sql.DB) ([]string, error) {
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

func offlinepeaple(online []string, allusers []string) []string {
	var ofline []string
	for _, y := range allusers {
		if notcontains(online, y) {
			ofline = append(ofline, y)
		}
	}

	return ofline
}

func notcontains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return false
		}
	}
	return true
}
