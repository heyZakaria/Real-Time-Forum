package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"forum/internal/app/models/utils"
	"forum/internal/app/websok"
	"maps"
	"net/http"
)

func GetOfflineUsersHandler(w http.ResponseWriter, r *http.Request) {

	onlinUsernames := websok.ChatHub.GetOnlineUsersnames()
	allUsers, err := allUsers(utils.Db1.Db)
	if err != nil {
		fmt.Println(err)
	}

	// this one is like this {zakaria:0, hassan:1}
	// to use this in frontend, in a loop add 0 to the first and 1 to the first ...
	// just add every element to the top offcourse if
	lastTalked, err := allUsersByLastMSG(utils.Db1.Db)
	if err != nil {
		fmt.Println(err)
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


func allUsersByLastMSG(db *sql.DB) (map[string]int, error) {
	query := `SELECT * FROM messages ORDER BY created_at `
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var lastTalked = make(map[string]int)
	//var users = []string{}
	var id int
	var sender string
	var receiver string
	var msg string
	var created_at string
	var count = 0
	for rows.Next() {

		if err := rows.Scan(&id, &sender, &receiver, &msg, &created_at); err != nil {
			fmt.Println("err scaning users ,user", err)
		}
		if _, ok := lastTalked[sender]; !ok {
			lastTalked[sender] = count
			count++
		} /* else if _, ok := lastTalked[receiver]; !ok {
			lastTalked[receiver] = count
			count++
		} */

		//fmt.Println("---------", id, sender, receiver, msg, created_at)
	}

	// fmt.Println("+++++++++", lastTalked)
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
	var offline = make(map[string]bool)
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
