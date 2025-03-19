package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"forum/internal/app/models/utils"
)

var (
	xhour    = 0
	xminut   = 0
	xseconde = 0
)

func AddPost(w http.ResponseWriter, r *http.Request) {
	var P utils.Posts
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowd", http.StatusMethodNotAllowed)
		return
	}

	json.NewDecoder(r.Body).Decode(&P)

	/// Trim space ??????????????
	/////////////////////////////////
	if len(P.Content) > 1000 || len(P.Title) > 100 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	cookie, err := r.Cookie("session_id")
	// fmt.Println("#$#$#", cookie)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	user_id, err := SelectUser(cookie.Value)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	if time.Now().Hour() == xhour && time.Now().Minute() == xminut && time.Now().Second()-xseconde < 5 {
		clearSession(w)
	} else {

		xhour = time.Now().Hour()
		xminut = time.Now().Minute()
		xseconde = time.Now().Second()

	}
	Creator, err := SelectUsername(user_id)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	P.Creator = Creator

	allKeys := make(map[int]bool)
	CategoryIDs := []int{}
	for _, categoryName := range P.Categories {

		CategoryID, _ := GetCategory(utils.Db1.Db, categoryName)
		if CategoryID != 0 {
			if _, ok := allKeys[CategoryID]; !ok {
				allKeys[CategoryID] = true
				CategoryIDs = append(CategoryIDs, CategoryID)
			}
		}

	}
	///////////////////////// check if category id
	post := strings.Trim(P.Content, " ")
	title := strings.Trim(P.Title, " ")
	if len(post) == 0 || len(title) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	postID, _ := InsertPost(utils.Db1.Db, P.Title, user_id, P.Content, P.Creator)
	P.Id = strconv.Itoa(postID)

	for _, categoryID := range CategoryIDs {
		err = AddCategoryPost(utils.Db1.Db, categoryID, postID)
		if err != nil {
			log.Fatal(err)
			return
		}
	}
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(P)
	P.Content = ""
	P.Title = ""
	P.Id = ""
	P.User_id = 0
}

func InsertPost(db *sql.DB, title string, user_id int, content string, creator string) (int, error) {
	query := "INSERT INTO posts (title, user_id,content,creator) VALUES (?,?,?,?)"

	res, err := db.Exec(query, title, user_id, content, creator)
	if err != nil {
		return 0, err
	}
	postID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(postID), err
}

func GetCategory(db *sql.DB, categoryName string) (int, error) {
	var categoryID int
	err := db.QueryRow("SELECT id FROM categories WHERE name_category = ?", categoryName).Scan(&categoryID)
	if err == sql.ErrNoRows {
		if err != nil {
			return 0, err
		}
	}

	return categoryID, nil
}

// I add post with his appropriate category
func AddCategoryPost(db *sql.DB, categoryID, postID int) error {
	query := "INSERT INTO post_categories(post_id,category_id) VALUES (?,?)"
	_, err := db.Exec(query, postID, categoryID)
	return err
}
