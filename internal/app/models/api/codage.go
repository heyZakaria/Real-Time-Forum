package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"forum/internal/app/models/utils"
)

func Codage(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "./mydatabase.db")
	if err != nil {
		log.Printf("Error opening database: %v", err)
		utils.MessageError(w, r, http.StatusInternalServerError ,"Database connection error")
		return

	}
	defer db.Close() 
	var posts []Posts

	posts, err = getPosts(db)
	if err != nil {
		log.Printf("Error fetching posts: %v", err)
		utils.MessageError(w, r, http.StatusInternalServerError, "error fetchin data")

	}

	likes_dislikes, err := GetLikesDislikes(db)
	if err != nil {
		log.Printf("Error fetching likes dislikes: %v", err)
		utils.MessageError(w, r, http.StatusInternalServerError, "error fetchin data")
	}

	comments, err := GetComments(db)
	if err != nil {
		log.Printf("Error fetching comments: %v", err)
		utils.MessageError(w, r, http.StatusInternalServerError, "error fetchin data")

	}

	comment_likes_dislikes, err := GetCommentLikesDislikes(db)

	if err != nil {
		log.Printf("Error fetching likes dislikes: %v", err)
		utils.MessageError(w, r, http.StatusInternalServerError, "error fetchin data")
	}

	categories, err := GetCategories(db)
	if err != nil {
		log.Printf("Error fetching Categories: %v", err)
		utils.MessageError(w, r, http.StatusInternalServerError, "error fetchin data")
	}
	postCategories, err := GetPostCategories(db)
	if err != nil {
		log.Printf("Error fetching post categories: %v", err)
		utils.MessageError(w, r, http.StatusInternalServerError, "error fetching data")
		return
	}

	NewPosts := GetCommentByID(comments, posts, comment_likes_dislikes)
	GetLikeAndDislike(NewPosts, likes_dislikes)

	AssignCategoriesToPosts(NewPosts, postCategories, categories)

	response := struct {
		NewPosts []Posts `json:"posts"`
	}{
		NewPosts: NewPosts,
	}
	// w.Header().Set("Content-Type", "application/json")

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ") // Add indentation for readability
	err = encoder.Encode(response)
	if err != nil {
		log.Printf("Error encoding JSON: %v", err)
		utils.MessageError(w, r, http.StatusInternalServerError,"Error creating JSON response")
		return
	}
}

func GetLikeAndDislike(Posts []Posts, like_dislike []Likes_dislikes) {
	for i := 0; i < len(Posts); i++ {
		for _, v := range like_dislike {
			if Posts[i].ID == v.Post_id {
				if v.Thetype == "LIKE" {

					Posts[i].Like++
					Posts[i].Likers = append(Posts[i].Likers, v.User_id)
					// likes_dislike = append(likes_dislike, like_dislike)

				} else {
					Posts[i].Dislike++

					Posts[i].Dislikers = append(Posts[i].Dislikers, v.User_id)

				}
			}
		}
	}
}

func GetComments(db *sql.DB) ([]Comments, error) {
	var comments []Comments
	query := "SELECT id, post_id,user_id,content FROM comments ORDER BY created_at DESC"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var comment Comments
		if err := rows.Scan(&comment.ID, &comment.Post_id, &comment.User_id, &comment.Content); err != nil {
			return nil, fmt.Errorf("scan error: %v", err)
		}

		comments = append(comments, comment)
	}
	return comments, nil
}

func getPosts(db *sql.DB) ([]Posts, error) {
	query := "SELECT id,title,content,user_id,creator FROM posts ORDER BY id DESC"
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("query error: %v", err)
	}
	defer rows.Close()

	var posts []Posts
	for rows.Next() {
		var post Posts

		if err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.User_id, &post.Creator); err != nil {
			return nil, fmt.Errorf("scan error: %v", err)
		}

		posts = append(posts, post)
	}

	return posts, nil
}
 
func GetCommentByID(comments []Comments, posts []Posts, comment_likes_dislikes []Comment_Likes_dislikes) []Posts {
	for i := 0; i < len(posts); i++ {
		for j := 0; j < len(comments); j++ {
			if posts[i].ID == comments[j].Post_id {
				for _, v := range comment_likes_dislikes {
					if comments[j].ID == v.Comment_Id {
						if v.Thetype == "LIKE" {
							comments[j].Like++
							comments[j].Likers = append(comments[j].Likers, v.User_id)

							} else {
								comments[j].Dislike++
								comments[j].Dislikers = append(comments[j].Dislikers, v.User_id)
								
							}
						}						
					}
					posts[i].Comment = append(posts[i].Comment, comments[j])
				}
		}
	}
	return posts
}



func GetLikesDislikes(db *sql.DB) ([]Likes_dislikes, error) {
	query := "SELECT post_id, thetype,user_id FROM likes_dislikes"
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("query error: %v", err)
	}
	defer rows.Close()

	var likes_dislike []Likes_dislikes // Use consistent variable name
	for rows.Next() {
		var like_dislike Likes_dislikes
		if err := rows.Scan(&like_dislike.Post_id, &like_dislike.Thetype, &like_dislike.User_id); err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}
		likes_dislike = append(likes_dislike, like_dislike) // Use the corrected variable name
	}

	// Check for any error that occurred during iteration
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during rows iteration: %v", err)
	}

	return likes_dislike, nil // Return the correct variable
}

func GetCommentLikesDislikes(db *sql.DB) ([]Comment_Likes_dislikes, error) {
	query := "SELECT comment_id, thetype, user_id FROM react_comments"
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("query error: %v", err)
	}
	defer rows.Close()

	var likes_dislike []Comment_Likes_dislikes // Use consistent variable name
	for rows.Next() {
		var like_dislike Comment_Likes_dislikes
		if err := rows.Scan(&like_dislike.Comment_Id, &like_dislike.Thetype, &like_dislike.User_id); err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}
		likes_dislike = append(likes_dislike, like_dislike) // Use the corrected variable name
	}

	// Check for any error that occurred during iteration
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during rows iteration: %v", err)
	}

	return likes_dislike, nil // Return the correct variable
}

func GetCategories(db *sql.DB) ([]Categories, error) {
	query := "SELECT id, name_category FROM categories"
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("query error: %v", err)
	}
	defer rows.Close()
	var Category []Categories
	for rows.Next() {
		var category Categories
		if err := rows.Scan(&category.ID, &category.Name_category); err != nil {
			return nil, fmt.Errorf("error scan %v", err)
		}
		Category = append(Category, category)
	}
	return Category, nil
}

func AssignCategoriesToPosts(posts []Posts, postCategories []Post_categories, categories []Categories) {
	categoryMap := make(map[int]string)
	for _, category := range categories {
		categoryMap[category.ID] = category.Name_category
	}

	for i, post := range posts {
		for _, pc := range postCategories {
			if post.ID == pc.Post_id {
				if categoryName, exists := categoryMap[pc.Category_id]; exists {
					posts[i].Categories = append(posts[i].Categories, categoryName)
				}
			}
		}
	}
}

func GetPostCategories(db *sql.DB) ([]Post_categories, error) {
	query := "SELECT post_id,category_id FROM post_categories"
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("query error: %v", err)
	}
	defer rows.Close()
	var Category []Post_categories
	for rows.Next() {
		var category Post_categories
		if err := rows.Scan(&category.Post_id, &category.Category_id); err != nil {
			return nil, fmt.Errorf("error scan %v", err)
		}
		Category = append(Category, category)
	}
	return Category, nil
}
