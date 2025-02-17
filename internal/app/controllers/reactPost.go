package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"forum/internal/app/models/utils"
)

func insertlike(db *sql.DB, post_id string, user_id int, Thetype string) error {
	query := "INSERT INTO likes_dislikes (post_id,user_id,Thetype) VALUES (?,?,?)"

	_, err := db.Exec(query, post_id, user_id, Thetype)
	return err
}

func insertdislike(db *sql.DB, post_id string, user_id int, Thetype string) error {
	query := "INSERT INTO likes_dislikes (post_id, user_id,Thetype) VALUES (?,?,?)"

	_, err := db.Exec(query, post_id, user_id, Thetype)
	return err
}

// remove like function is just delet from likes_dislikes WHERE post_id so all you need is post id
func removelike(db *sql.DB, user_id int, post_id string) error {
	query := "DELETE FROM likes_dislikes WHERE post_id = ? AND user_id = ? AND Thetype = 'LIKE'"
	_, err := db.Exec(query, post_id, user_id)
	return err
}

/// same same but defrente line

// remove like function is just delet from likes_dislikes WHERE post_id so all you need is post id
func removedislike(db *sql.DB, user_id int, post_id string) error {
	query := "DELETE FROM likes_dislikes WHERE post_id = ? AND user_id AND Thetype = 'DISLIKE'"
	_, err := db.Exec(query, post_id, user_id)
	return err
}

/// same same but defrente line

func getuserid(r *http.Request) int {
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

func fetchLikeCount(db *sql.DB, post_id string) (int, error) {
	query := "SELECT COUNT(*) FROM likes_dislikes WHERE post_id = ? AND Thetype = 'LIKE'"
	var count int
	err := db.QueryRow(query, post_id).Scan(&count)
	return count, err
}

func fetchDislikeCount(db *sql.DB, post_id string) (int, error) {
	query := "SELECT COUNT(*) FROM likes_dislikes WHERE post_id = ? AND Thetype = 'DISLIKE'"
	var count int
	err := db.QueryRow(query, post_id).Scan(&count)
	return count, err
}

type ReactionRequest struct {
	Action string `json:"action"` // "increment" olaaa "decrement"
}

type ReactionResponse struct {
	Success  bool `json:"success"`
	Likes    int  `json:"likes"`
	Dislikes int  `json:"dislikes"`
}

// get post from url dik /api/posts/{id}/like from fetch.js
func getPostIDFromURL(path string) string {
	//  /api/posts/{id}/like
	parts := strings.Split(path, "/")
	if len(parts) >= 4 {
		return parts[3] // returns   ID
	}
	return ""
}

func getreaction(path string) string {
	parts := strings.Split(path, "/")
	if len(parts) >= 4 {
		return parts[4] // returns   reaction
	}
	return ""
}

func HandleReaction(w http.ResponseWriter, r *http.Request) {
	user_id := getuserid(r)
	if user_id < 1 {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	reaction := getreaction(r.URL.Path)

	if r.Method != "POST" { // pritect liknk /api/posts/{id}/like
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if reaction == "like" {

		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		postID := getPostIDFromURL(r.URL.Path)
		if postID == "" {
			http.Error(w, "Invalid post ID", http.StatusBadRequest)
			return
		}

		var req ReactionRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		/// Before updating check if the ACTION IS EXISTS

		// Update like count in database based on action
		var newLikeCount int
		if req.Action == "increment" {
			err0 := removelike(utils.Db1.Db, user_id, postID)
			err := insertlike(utils.Db1.Db, postID, user_id, "LIKE")

			if err != nil || err0 != nil {
				w.WriteHeader(500)
				// utils.MessageError(w, r, http.StatusMethodNotAllowed, err0.Error())
				return

			}

		} else if req.Action == "decrement" {

			err0 := removelike(utils.Db1.Db, user_id, postID)

			if err0 != nil {
				w.WriteHeader(500)
				return

			}

		} else {
			http.Error(w, "Invalid action", http.StatusBadRequest)
			return
		}

		//  ila bghiti t3rf x7al dyal like wlaw

		newLikeCount, err := fetchLikeCount(utils.Db1.Db, postID)
		if err != nil {
			http.Error(w, "Failed to fetch like count", http.StatusInternalServerError)
			return
		}

		newDislikeCount, err := fetchDislikeCount(utils.Db1.Db, postID)
		if err != nil {
			http.Error(w, "Failed to fetch dislike count", http.StatusInternalServerError)
			return
		}

		response := ReactionResponse{
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

		postID := getPostIDFromURL(r.URL.Path)
		if postID == "" {
			http.Error(w, "Invalid post ID", http.StatusBadRequest)
			return
		}

		var req ReactionRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Update like count in database based on action
		var newLikeCount int
		if req.Action == "increment" {
			err0 := removedislike(utils.Db1.Db, user_id, postID)
			err := insertdislike(utils.Db1.Db, postID, user_id, "DISLIKE")

			if err0 != nil || err != nil {
				http.Error(w, "Failed to fetch like count", http.StatusInternalServerError)

				return

			}

		} else if req.Action == "decrement" {

			err0 := removedislike(utils.Db1.Db, user_id, postID)

			if err0 != nil {
				http.Error(w, "Failed to fetch like count", http.StatusInternalServerError)
				return

			}

		} else {
			http.Error(w, "Invalid action", http.StatusBadRequest)
			return
		}

		newDislikeCount, err := fetchDislikeCount(utils.Db1.Db, postID)
		if err != nil {
			http.Error(w, "Failed to fetch like count", http.StatusInternalServerError)
			return
		}

		response := ReactionResponse{
			Success:  true,
			Likes:    newLikeCount,
			Dislikes: newDislikeCount,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)

	}
}
