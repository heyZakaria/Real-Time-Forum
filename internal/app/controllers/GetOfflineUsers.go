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

type LMS websok.Storeactivewebsocketclient

func GetOfflineUsersHandler(w http.ResponseWriter, r *http.Request) {

	onlinUsernames := websok.ChatHub.GetOnlineUsersnames()
	allUsers, err := allUsers(utils.Db1.Db)
	if err != nil {
		fmt.Println(err)
	}

	_, err = allUsersByLastMSG(utils.Db1.Db)
	if err != nil {
		fmt.Println(err)
	}

	offlineUsers := offlinePeople(onlinUsernames, allUsers)

	// This copies all onlineusername to offlineUsers
	// so we have one object, I need sorting
	maps.Copy(offlineUsers, onlinUsernames)
	//fmt.Println(websok.ChatHub.LastInertedMsg)
	//lastChatUsers(utils.Db1.Db, int(websok.ChatHub.LastInertedMsg))
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(offlineUsers)
}

func lastChatUsers(db *sql.DB, LIM int) {
	query := `SELECT * FROM messages WHERE id = ?`
	row := db.QueryRow(query, LIM)
	var data string
	row.Scan(data)
	fmt.Println(data)
}
func allUsersByLastMSG(db *sql.DB) ([]string, error) {
	query := `SELECT * FROM messages ORDER BY created_at`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []string
	for rows.Next() {
		var id int
		var sender string
		var receiver string
		var msg string
		var created_at string

		if err := rows.Scan(&id,&sender, &receiver, &msg, &created_at ); err != nil {
			fmt.Println("err scaning users ,user", err)
		}
		users = append(users, sender, receiver)
		
		fmt.Println("---------", id, sender, receiver, msg, created_at )
	}
	return users, nil
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
