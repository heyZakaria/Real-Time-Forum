package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"regexp"
	"strings"
	"unicode"

	"forum/internal/app/models/utils"

	"golang.org/x/crypto/bcrypt"
)

type Request struct {
	Username   string `json:"username"`
	Age        string `json:"age"`
	Gender     string `json:"genderIs"`
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}

var (
	validusername = true
	validemail    = true
	validpassword = true
)

func Registration(w http.ResponseWriter, r *http.Request) {

	page := []string{"internal/app/views/templates/forum.html"}
	if r.Method == http.MethodGet {
		utils.ExecuteTemplate(w, page, nil)

	} else if r.Method == http.MethodPost {

		request := Request{}
		response := make(map[string]bool)
		// Décoder le corps de la requête JSON
		err := json.NewDecoder(r.Body).Decode(&request)

		if err != nil {
			response["isValidata"] = false
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
			return
		}

		// checkusername
		if !Printable(request.Username) || len(request.Username) > 25 || len(request.Username) < 2 || strings.TrimSpace(request.Username) == "" {
			validusername = false
		}
		// checkemail
		if !isValidEmail(request.Email) || len(request.Email) > 60 || len(request.Email) < 8 || !Printable(request.Email) || strings.TrimSpace(request.Email) == "" {
			validemail = false
		}
		// checkpassword
		if len(request.Password) < 8 || len(request.Password) > 50 ||
			!Printable(request.Password) || !CheckCaractere(request.Password) {
			validpassword = false
		}

		if !validemail || !validpassword || !validusername {
			response["isValiddata"] = false
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
			return
		}

		exist, err := EmailOrUsernameExiste(utils.Db1.Db, request.Email, request.Username)
		if exist && err == nil {
			response["emilorusernameexsist"] = true
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
			return

		} else if err != nil {
			response["InternalError"] = true
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(response)
			return
		} else if !exist {
			hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
			var gender = ""
			if request.Gender == "F" {
				gender = "F"
			} else {
				gender = "M"

			}
			username := strings.ToLower(request.Username)
			err := insereData(utils.Db1.Db, request.Email, request.First_name, request.Last_name, gender, username, string(hashedPassword), request.Age)
			if err != nil {
				response["InternalError"] = true
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(response)
				return
			} else {

				response["emilorusernameexsist"] = false
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(response)
				return
			}
		}
	}
}

func isValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

func Printable(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] < 32 {
			return false
		}
	}
	return true
}

func CheckCaractere(s string) bool {
	hasUpper := false
	hasLower := false
	hasDigit := false
	hasSpecial := false

	for _, ch := range s {
		if unicode.IsUpper(ch) {
			hasUpper = true
		} else if unicode.IsLower(ch) {
			hasLower = true
		} else if unicode.IsDigit(ch) {
			hasDigit = true
		} else if !unicode.IsLetter(ch) && !unicode.IsDigit(ch) {
			hasSpecial = true
		}
	}

	return hasUpper && hasLower && hasDigit && hasSpecial
}

func EmailOrUsernameExiste(db *sql.DB, email, username string) (bool, error) {
	query := "SELECT COUNT(*) FROM users WHERE email = ? OR username = ?"

	var count int
	err := db.QueryRow(query, email, username).Scan(&count)
	if err != nil {
		return false, err
	}

	if count > 0 {
		return true, nil
	}
	return false, nil
}

func insereData(db *sql.DB, email, first_name, last_name, gender, username, password, age string) error {
	query := "INSERT INTO users (username,age, gender,first_name,last_name,  email, password_hash) VALUES (?, ?, ?, ?, ?, ?,?)"

	_, err := db.Exec(query, username, age, gender, first_name, last_name, email, password)
	return err
}
