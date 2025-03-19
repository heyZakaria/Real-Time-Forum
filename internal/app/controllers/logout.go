package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"forum/internal/app/models/utils"
	"forum/internal/app/websok"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.MessageError(w, r, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	realtoken := r.URL.Query().Get("token")
	token := realtoken[11:]

	//fmt.Println("tooken kamla", token[11:])

	//fmt.Println("tooken", token)

	// token, err := r.Cookie("session_id")
	// // fmt.Println(token,"@@@@@@@@@@@@@@@@@")
	// if err != nil {
	// 	fmt.Println("11")
	// }
	user_id, err := SelectUser(token)
	//fmt.Println("useridlogu", user_id)
	if err != nil {
		fmt.Println("22", err)
	}
	username, err := SelectUsername(user_id)
	//fmt.Println("usernamelogout:", username)

	if err != nil {
		fmt.Println("33", err)
	}

	// RemoveSessionFromDB(utils.Db1.Db, token)
	clearSession(w)

	//fmt.Println("REMOEV")
	// delete(h.Clients, c.Username)
	// c.Conn.Close()

	// fmt.Println(websok.UnregisterClientByUsername("xxxx"))
	if websok.UnregisterClientByUsername(username) {
		fmt.Println("closse for that:", username)
	}

	json.NewEncoder(w).Encode("Success")
}

func clearSession(w http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "session_id",
		Value:  "",
		Path:   "/",
		MaxAge: 0,
	}
	http.SetCookie(w, cookie)
}

func RemoveSessionFromDB(db *sql.DB, token *http.Cookie) error {
	query := "DELETE  FROM session WHERE code = ?"
	_, err := db.Exec(query, token)
	return err
}
