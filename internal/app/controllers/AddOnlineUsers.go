package controllers

import (
	"database/sql"
	"fmt"
)

func DeleteUserFromOnlineUsers(db *sql.DB, userID int) error {
	_, err := db.Exec("DELETE FROM online_users WHERE user_id = ?", userID)
	if err != nil {
		fmt.Println("err deleting user from online users ", err)
		return err
	}
	return nil
}


 