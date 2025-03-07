package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

 	"forum/internal/app/models/utils"

	"golang.org/x/crypto/bcrypt"
)

var (
	isEmail = false
	Error   = false
)

type Requiste struct {
	Emailorusername string `json:"emailorusername"` // Identifier
	Password        string `json:"password"`
}

var sessions = make(map[string]utils.Session) // session 7athoum f database fi 7alat ila server tfa o ch3el ghadi dima tlqa data tma ama iladrtiha hna maghdich tlqaha

// Helper to generate random session IDs

type Session struct {
	Username  string
	ExpiresAt time.Time 
}

// Helper to generate random session IDs
func generateSessionID() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	sessionID := make([]byte, 32)
	for i := range sessionID {
		sessionID[i] = charset[rand.Intn(len(charset))]
	}
	return string(sessionID)
}

func Login(w http.ResponseWriter, r *http.Request) {
	page := []string{"internal/app/views/templates/forum.html"}
	fmt.Println(r.URL.Path, "HEEERE")

	var G utils.User
	if r.Method == http.MethodGet {
		utils.ExecuteTemplate(w, page, nil)
	} else if r.Method == http.MethodPost {

		// w.Header().Set("Content-Type", "application/json")

		request := Requiste{}
		response := make(map[string]bool)

		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {

			response["isValidData"] = false
			// why badRequest just invalid infos
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(response)
			return

		}
		if len(request.Emailorusername) < 2 || len(request.Emailorusername) > 60 || len(request.Password) < 8 || len(request.Password) > 50 || !Printable(request.Emailorusername) || !Printable(request.Password) {
			response["isValidData"] = false

			// why badRequest just invalid infos
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(response)
			return
		}

		isAuthneticated, err := authenticateUser(utils.Db1.Db, request.Emailorusername, request.Password)
		if err != nil {
			response["isValidData"] = false

			// toooo
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(response)
			return
		}
		if isAuthneticated {

			//   new session
			//???????????????????????????????
			/// Wach USERNAME wla EMAIL and WHY
			id := CheckIsEmailOrUsername(request.Emailorusername, utils.Db1.Db)

			if Error || id == 0 {
				response["errserver"] = true
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(response)
				return

			} else if isEmail {
				G.Username = Getusername(request.Emailorusername, utils.Db1.Db)
			} else {
				G.Username = request.Emailorusername
			}
			G.ID = id

			err := DeleteIfAnySession(utils.Db1.Db, id)

			if err {
				response["errserver"] = true
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(response)
			} else {

				sessionID := generateSessionID()
				expiration := time.Now().Add(30 * time.Minute) // S e s s i o n 30 minutes

				// Add user to online_users table


				sessions[sessionID] = utils.Session{
					Username:  G.Username,
					ExpiresAt: expiration,
				}

				cookie := &http.Cookie{
					Name:    "session_id",
					Value:   sessionID,
					Expires: expiration,
				}

				http.SetCookie(w, cookie)
			

				// add user online statu 

				err1 :=AddUserToOnlineUsers(utils.Db1.Db,G.ID,G.Username)
				if err1 != nil {
					utils.MessageError(w, r, http.StatusInternalServerError, "Error update online statu")
					return
				}
			














				////////////////////////

				err := AddUserToDatabase(sessionID, G.ID, utils.Db1.Db)
				if err != nil {
					response["errserver"] = true
					w.WriteHeader(http.StatusInternalServerError)
					json.NewEncoder(w).Encode(response)
					return

				}

				response["isValidData"] = true
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(response)
			}

		} //????????????????????????????
		// ELSE
	}
}

func AddUserToDatabase(sessionID string, ID int, DB *sql.DB) error {
	query := "INSERT INTO session (id_users, code) VALUES (?, ?)"

	_, err := DB.Exec(query, ID, sessionID)

	return err
}

func DeleteIfAnySession(db *sql.DB, id int) bool {
	query := "DELETE FROM session WHERE id_users = ?"

	_, err := db.Exec(query, id)
	if err != nil {
		return true
	}

	return false
}

func Getusername(email string, DB *sql.DB) string {
	var username string
	query := "SELECT username FROM users WHERE email = ?"

	row := DB.QueryRow(query, email)

	err := row.Scan(&username)
	if err != nil {
		return "Error"
	}

	return username
}

func CheckIsEmailOrUsername(EmailOrUsername string, DB *sql.DB) int {
	var id int

	quirie := "SELECT id FROM users WHERE username = ? OR email = ?"
	row := DB.QueryRow(quirie, EmailOrUsername, EmailOrUsername)
	err := row.Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			isEmail = true
			return 0
		} else {
			Error = true
			return 0
		}
	}
	return id
}

func GetId(username string, DB *sql.DB) (int, error) {
	var id int
	quirie := "SELECT id FROM users WHERE username  = ? "
	row := DB.QueryRow(quirie, username)
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func checkPassword(storedPassword, enteredPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(enteredPassword))
	return err == nil
}

func authenticateUser(db *sql.DB, EmailOrUsername, enteredPassword string) (bool, error) {
	var storedPassword string
	query := "SELECT password_hash FROM users WHERE email= ? OR username = ?"
	err := db.QueryRow(query, EmailOrUsername, EmailOrUsername).Scan(&storedPassword)
	if err != nil {
		return false, err
	}

	if checkPassword(storedPassword, enteredPassword) {
		return true, nil
	} else {
		return false, fmt.Errorf("invalid password")
	}
}



func AddUserToOnlineUsers(db *sql.DB, userID int,username string) error {
	//need to add ON CONFLICT(user_id) DO UPDATE SET last_active = ? 



	_, err := db.Exec("INSERT INTO online_users (user_id, last_active,username) VALUES (?, ?,?) ", userID,time.Now(),username)
	if err != nil {
		fmt.Println("err add user on online users :", err)
		return err
	}
	return nil
}
