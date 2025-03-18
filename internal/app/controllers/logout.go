package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"forum/internal/app/models/utils"
	"forum/internal/app/websok"
	"net/http"
)

var c websok.Client
var h websok.Storeactivewebsocketclient

func Logout(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.MessageError(w, r, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}

	token, _ := r.Cookie("session_id")

	RemoveSessionFromDB(utils.Db1.Db, token)
	clearSession(w)
	fmt.Println( "REMOEV")
	//delete(h.Clients, c.Username)

	// c.Conn.Close()
	json.NewEncoder(w).Encode("Success")

}

func clearSession(w http.ResponseWriter) {

	cookie := &http.Cookie{
		Name:   "session_id",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)

}

func RemoveSessionFromDB(db *sql.DB, token *http.Cookie) error {
	query := "DELETE  FROM session WHERE code = ?"
	_, err := db.Exec(query, token)
	return err
}
