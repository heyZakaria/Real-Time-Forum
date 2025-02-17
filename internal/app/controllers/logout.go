package controllers

import (
	"database/sql"
	"encoding/json"
	"forum/internal/app/models/utils"
	"net/http"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.MessageError(w, r, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}

	token, _ := r.Cookie("session_id")

	RemoveSessionFromDB(utils.Db1.Db, token)
	clearSession(w)
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
