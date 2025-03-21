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

	
	user_id, err := SelectUser(token)
	if err != nil {
		http.Error(w, "username not fund", http.StatusInternalServerError)
	}
	username, err := SelectUsername(user_id)

	if err != nil {
		http.Error(w, "username not fund", http.StatusInternalServerError)
	}

	// RemoveSessionFromDB(utils.Db1.Db, token)
	clearSession(w)

	
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
