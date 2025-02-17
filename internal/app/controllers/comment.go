package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"forum/internal/app/models/utils"
)

func AddComment(w http.ResponseWriter, r *http.Request) {
	var C utils.Comment
	if r.Method != http.MethodPost {
		utils.MessageError(w, r, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}

	json.NewDecoder(r.Body).Decode(&C)
	if len(C.Content) > 500 {
		utils.MessageError(w, r, http.StatusBadRequest, "Bad Req")
		return
	}
	cookie, err := r.Cookie("session_id")
	if err != nil {
		utils.MessageError(w, r, http.StatusUnauthorized, "Siir f7aleeek")
		return
	}
	user_id, err := SelectUser(cookie.Value)
	if err != nil {
		utils.MessageError(w, r, http.StatusUnauthorized, "Siir f7aleeek")
		return
	}
	_, err = SelectUsername(user_id)
	if err != nil {
		utils.MessageError(w, r, http.StatusUnauthorized, "Siir f7aleeek")
		return
	}
	comment := strings.Trim(C.Content, " ")

	if len(comment) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	commentID, _ := insertComment(utils.Db1.Db, C.Post_id, user_id, C.Content)
	C.Id = strconv.Itoa(commentID)
	json.NewEncoder(w).Encode(C)

}

func insertComment(db *sql.DB, post_id string, user_id int, content string) (int, error) {
	query := "INSERT INTO comments (post_id, user_id, content) VALUES (?, ?, ?)"

	res, _ := db.Exec(query, post_id, user_id, content)

	commentID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(commentID), err
}
