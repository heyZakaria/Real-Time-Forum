package utils

import (
	"database/sql"
	"time"
)

type Posts struct {
	Id         string   `json:"id"`
	Title      string   `json:"title"`
	Content    string   `json:"content"`
	User_id    int      `json:"user_id"`
	Creator    string   `json:"creator"`
	Categories []string `json:"category"`
}
type User struct {
	ID       int
	Email    string
	Password string
	Username string
}
type Comment struct {
	Id        string `json:"id"`
	Post_id   string `json:"post_id"`
	Content   string `json:"content"`
	User_id   int    `json:"user_id"`
	Creator   string `json:"creator"`
	Like      int    `json:"like"`
	Dislike   int    `json:"dislike"`
	Likers    []int  `json:"likers"`
	Dislikers []int  `json:"dislikers"`
}

type Info struct {
	Username  string
	Email     string
	Password  string
	Password2 string
}
type Db struct {
	Db *sql.DB
}

var Db1 Db

type Session struct {
	Username  string
	ExpiresAt time.Time
}

var Session1 Session
