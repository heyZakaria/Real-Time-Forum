package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"forum/internal/app/models/utils"
)

func Home(w http.ResponseWriter, r *http.Request) {
	page := []string{"internal/app/views/templates/forum.html"}
	if r.URL.Path != "/register" && r.URL.Path != "/login" && r.URL.Path != "/" {
		utils.MessageError(w, r, http.StatusNotFound, "Page Not Found!!!")
		return
	}
	if r.URL.Path == "/register" {

		// utils.MessageError(w, r, http.StatusNotFound, "Page Not Found!!!!")
		Registration(w, r)
		return
	}

	if r.URL.Path == "/login" {
		// utils.MessageError(w, r, http.StatusNotFound, "Page Not Found!!!!")
		Login(w, r)
		return
	}

	if r.Method == "POST" {

		utils.ExecuteTemplate(w, page, "")

	} else if r.Method == "GET" {
		utils.ExecuteTemplate(w, page, "")
		return
	}
}

func CheckRegistration(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		utils.MessageError(w, r, http.StatusUnauthorized, "Access denied")
		return
	}
	if CheckUserInDB(cookie.Value) {

		user_id, err := SelectUser(cookie.Value)
		if err != nil {
			utils.MessageError(w, r, http.StatusInternalServerError, "Internal Server Error")
			return
		}

		w.WriteHeader(200)
		json.NewEncoder(w).Encode(user_id)

	} else {
		utils.MessageError(w, r, http.StatusUnauthorized, "Access denied")
		return
	}
}

func CheckUserInDB(token string) bool {
	query := "SELECT id_users FROM session WHERE code= ?"
	hold := 0
	utils.Db1.Db.QueryRow(query, token).Scan(&hold)

	if hold != 0 {
		return hold != 0
	}

	return false
}

func SelectUser(token string) (int, error) {

	id_user := 0
	query := "SELECT id_users FROM session WHERE code= ?"
	err := utils.Db1.Db.QueryRow(query, token).Scan(&id_user)
	if err != nil {
		if err == sql.ErrNoRows {

			return 0, err
		}
		return 0, err
	}

	return id_user, nil
}

func SelectUsername(id int) (string, error) {
	var P utils.Posts

	query1 := "SELECT username FROM users WHERE id = ?"
	err := utils.Db1.Db.QueryRow(query1, id).Scan(&P.Creator)
	if err != nil {
		if err == sql.ErrNoRows {

			return "", err
		}
		return "", err
	}
	return P.Creator, nil
}
