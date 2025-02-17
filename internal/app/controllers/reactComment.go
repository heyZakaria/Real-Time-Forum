package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"forum/internal/app/models/utils"
)

func insertCommentLike(db *sql.DB, comment_Id string, user_id int, Thetype string) error {
	query := "INSERT INTO react_comments (comment_Id,user_id,Thetype) VALUES (?,?,?)"

	_, err := db.Exec(query, comment_Id, user_id, Thetype)
	return err
}

func insertCommentDislike(db *sql.DB, comment_Id string, user_id int, Thetype string) error {
	query := "INSERT INTO react_comments (comment_Id, user_id,Thetype) VALUES (?,?,?)"

	_, err := db.Exec(query, comment_Id, user_id, Thetype)
	return err
}

// remove like function is just delet from likes_dislikes WHERE comment_Id so all you need is post id
func removeCommentLike(db *sql.DB, user_id int, comment_Id string) error {
	query := "DELETE FROM react_comments WHERE comment_Id = ? AND user_id = ? AND Thetype = 'LIKE'"
	_, err := db.Exec(query, comment_Id, user_id)
	return err
}

/// same same but defrente line

// remove like function is just delet from likes_dislikes WHERE comment_Id so all you need is post id
func removeCommentDislike(db *sql.DB, user_id int, comment_Id string) error {
	query := "DELETE FROM react_comments WHERE comment_Id = ? AND user_id AND Thetype = 'DISLIKE'"
	_, err := db.Exec(query, comment_Id, user_id)
	return err
}

/// same same but defrente line

func getUserID(r *http.Request) int {
	id_user := 1
	cookie, err := r.Cookie("session_id")
	if err != nil {
		return 0
	}
	query := "SELECT id_users FROM session WHERE code= ?"
	err = utils.Db1.Db.QueryRow(query, cookie.Value).Scan(&id_user)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Errorf("utilisateur non trouvé")
		}
		fmt.Errorf("xi 7aja non trouvé")

	}

	return id_user
}

func fetchCommentLikeCount(db *sql.DB, comment_Id string) (int, error) {
	query := "SELECT COUNT(*) FROM react_comments WHERE comment_Id = ? AND Thetype = 'LIKE'"
	var count int
	err := db.QueryRow(query, comment_Id).Scan(&count)
	return count, err
}

func fetchCommentDislikeCount(db *sql.DB, comment_Id string) (int, error) {
	query := "SELECT COUNT(*) FROM react_comments WHERE comment_Id = ? AND Thetype = 'DISLIKE'"
	var count int
	err := db.QueryRow(query, comment_Id).Scan(&count)
	return count, err
}

type CommentLikeRequest struct {
	Action string `json:"action"` // "increment" olaaa "decrement"
}

type CommentLikeResponse struct {
	Success  bool `json:"success"`
	Likes    int  `json:"likes"`
	Dislikes int  `json:"dislikes"`
}

type CommentDislikeRequest struct {
	Action string `json:"action"` // "increment" olaaa "decrement"
}

type CommentDislikeResponse struct {
	Success bool `json:"success"`
	Likes   int  `json:"likes"`
}

// get post from url dik /api/posts/{id}/like from fetch.js
func getCommentIDFromURL(path string) string {
	//  /api/posts/{id}/like
	parts := strings.Split(path, "/")
	if len(parts) >= 4 {
		return parts[3] // returns   ID
	}
	return ""
}

func getCommentReaction(path string) string {
	parts := strings.Split(path, "/")
	if len(parts) >= 4 {
		return parts[4] // returns   reaction
	}
	return ""
}

func HandleCommentReaction(w http.ResponseWriter, r *http.Request) {
	user_id := getUserID(r)
	if user_id < 1 {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	reaction := getCommentReaction(r.URL.Path)

	if reaction == "like" {

		if r.Method != "POST" { // pritect liknk /api/posts/{id}/like
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		commentID := getCommentIDFromURL(r.URL.Path)
		if commentID == "" {
			http.Error(w, "Invalid post ID", http.StatusBadRequest)
			return
		}

		var req CommentLikeRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Update like count in database based on action
		var newLikeCount int
		if req.Action == "increment" {
			err0 := removeCommentLike(utils.Db1.Db, user_id, commentID)
			err := insertCommentLike(utils.Db1.Db, commentID, user_id, "LIKE")

			if err != nil || err0 != nil {
				http.Error(w, "Failed to fetch like count", http.StatusInternalServerError)

				return

			}

		} else if req.Action == "decrement" {

			err0 := removeCommentLike(utils.Db1.Db, user_id, commentID)

			if err0 != nil {
				http.Error(w, "Failed to fetch like count", http.StatusInternalServerError)
				return

			}

		} else {
			http.Error(w, "Invalid action", http.StatusBadRequest)
			return
		}

		//  ila bghiti t3rf x7al dyal like wlaw

		newLikeCount, err := fetchCommentLikeCount(utils.Db1.Db, commentID)
		if err != nil {

			http.Error(w, "Failed to fetch like count", http.StatusInternalServerError)
			return
		}

		newDislikeCount, err := fetchCommentDislikeCount(utils.Db1.Db, commentID)
		if err != nil {

			http.Error(w, "Failed to fetch dislike count", http.StatusInternalServerError)
			return
		}

		response := CommentLikeResponse{
			Success:  true,
			Likes:    newLikeCount,
			Dislikes: newDislikeCount,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)

	} else if reaction == "dislike" {

		if r.Method != "POST" { // pritect liknk /api/posts/{id}/like
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		commentID := getCommentIDFromURL(r.URL.Path)
		if commentID == "" {
			http.Error(w, "Invalid post ID", http.StatusBadRequest)
			return
		}

		var req CommentLikeRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Update like count in database based on action
		var newLikeCount int
		if req.Action == "increment" {
			err := removeCommentDislike(utils.Db1.Db, user_id, commentID)

			err0 := insertCommentDislike(utils.Db1.Db, commentID, user_id, "DISLIKE")

			if err0 != nil || err != nil {
				http.Error(w, "Failed to fetch like count", http.StatusInternalServerError)
				return

			}

		} else if req.Action == "decrement" {

			err0 := removeCommentDislike(utils.Db1.Db, user_id, commentID)

			if err0 != nil {
				http.Error(w, "Failed to fetch like count", http.StatusInternalServerError)

				return

			}

		} else {
			http.Error(w, "Invalid action", http.StatusBadRequest)
			return
		}

		newDislikeCount, err := fetchCommentDislikeCount(utils.Db1.Db, commentID)
		if err != nil {

			http.Error(w, "Failed to fetch like count", http.StatusInternalServerError)
			return
		}

		response := CommentLikeResponse{
			Success:  true,
			Likes:    newLikeCount,
			Dislikes: newDislikeCount,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)

	}
}
