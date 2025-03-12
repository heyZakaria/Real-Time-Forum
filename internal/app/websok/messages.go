package websok

import (
	"database/sql"
	"time"
)

// Message represents a simple message chat
type Message struct {
	ID         int64     `json:"id"`
	SenderID   string    `json:"sender_id"`
	ReceiverID string    `json:"receiver_id"`
	Content    string    `json:"message_content"`
	CreatedAt  time.Time `json:"created_at"`
}
func SaveMessage(db *sql.DB, senderID, receiverID, content string) (int64, error) {
	query := `INSERT INTO messages (sender_id, receiver_id, message_content, created_at) VALUES (?, ?, ?, CURRENT_TIMESTAMP)`

	result, err := db.Exec(query, senderID, receiverID, content)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func GetConversationHistory(db *sql.DB, userID1, userID2 string, limit int) ([]Message, error) {
	query := `
	SELECT id, sender_id, receiver_id, message_content, created_at
	FROM messages
	WHERE (sender_id = ? AND receiver_id = ?) OR (sender_id = ? AND receiver_id = ?)
	ORDER BY created_at ASC
	LIMIT ?`

	rows, err := db.Query(query, userID1, userID2, userID2, userID1, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []Message
	for rows.Next() {
		var msg Message
		if err := rows.Scan(&msg.ID, &msg.SenderID, &msg.ReceiverID, &msg.Content, &msg.CreatedAt); err != nil {
			return nil, err
		}
		messages = append(messages, msg)
	}

	return messages, nil
}
