package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"forum/internal/app/models/utils"
)

/////// next will go in codage api

func GetOnlineUsers(w http.ResponseWriter, r *http.Request) {
	rows, err := utils.Db1.Db.Query(`
        SELECT users.id, users.username 
        FROM online_users 
        JOIN users ON online_users.user_id = users.id
    `)
	if err != nil {
		http.Error(w, "Error fetching online users", http.StatusInternalServerError)
		fmt.Println("Error fetching online users:", err)
		return
	}
	defer rows.Close()

	var users []map[string]interface{}

	for rows.Next() {
		var id int
		var username string
		if err := rows.Scan(&id, &username); err != nil {
			http.Error(w, "Error reading data", http.StatusInternalServerError)
			return
		}
		users = append(users, map[string]interface{}{
			"id":       id,
			"username": username,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
