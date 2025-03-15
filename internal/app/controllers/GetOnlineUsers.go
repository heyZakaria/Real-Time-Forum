package controllers

import (
	"encoding/json"
	"forum/internal/app/websok"
	"net/http"
)

/////// next will go in codage api

func GetOnlineUsersHandler(w http.ResponseWriter, r *http.Request) {
	//get onlin users from websocket
	usernames := websok.ChatHub.GetOnlineUsersnames()

	//my struct to the response
	type UserResponse struct {
		Username string `json:"username"`
	}
	response := make([]UserResponse, len(usernames))

	for i, username := range usernames {
		response[i] = UserResponse{
			Username: username,
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}
