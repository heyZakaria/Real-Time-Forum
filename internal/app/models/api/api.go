package api

import "time"

type Posts struct {
	ID         int        `json:"id"`
	User_id    int        `json:"user_id"`
	Title      string     `json:"title"`
	Content    string     `json:"content"`
	Comment    []Comments `json:"comment"`
	Like       int        `json:"like"`
	Dislike    int        `json:"dislike"`
	Creator    string     `json:"creator"`
	Likers     []int      `json:"likers"`
	Dislikers  []int      `json:"dislikers"`
	Categories []string   `json:"categories"`
}

type Likes_dislikes struct {
	// ID      int    `json:"id"`
	Post_id int    `json:"post_id"`
	User_id int    `json:"user_id"`
	Thetype string `json:"thetype"` // ikhan maxi 'type'
}

type Comments struct {
	ID         int       `json:"id"`
	Content    string    `json:"content"`
	Post_id    int       `json:"post_id"`
	User_id    int       `json:"user_id"`
	Like       int       `json:"like"`
	Dislike    int       `json:"dislike"`
	Likers     []int     `json:"likers"`
	Dislikers  []int     `json:"dislikers"`
	Created_at time.Time `json:"created_at"`
}

type Comment_Likes_dislikes struct {
	ID         int    `json:"id"`
	Comment_Id int    `json:"comment_id"`
	User_id    int    `json:"user_id"`
	Thetype    string `json:"thetype"` // ikhan maxi 'type'
}

type Post_categories struct {
	Post_id     int `json:"post_id"`
	Category_id int `json:"category_id"`
}

type Categories struct {
	ID            int    `json:"id"`
	Name_category string `json:"name_category"`
}
