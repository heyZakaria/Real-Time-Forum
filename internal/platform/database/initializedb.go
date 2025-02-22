package database

import (
	"database/sql"
	"log"
	"os"
)

func CreateDatabase() (*sql.DB, error) {
	var db *sql.DB
	db, err := sql.Open("sqlite3", "./mydatabase.db")
	if err != nil {
		log.Fatal(err)
	}

	schema, err := os.ReadFile("./internal/platform/database/schema.sql")
	if err != nil {
		log.Fatal("err1", err)
	}

	_, err = db.Exec(string(schema))
	if err != nil {
		log.Fatal("err", err)
	}

	creatCategory(db, "sport")
	creatCategory(db, "science")
	creatCategory(db, "entertainment")

	return db, nil
}

func creatCategory(db *sql.DB, categoryName string) {
	var categoryID int

	_ = db.QueryRow("SELECT id FROM categories WHERE name_category = ?", categoryName).Scan(&categoryID)

	if categoryID == 0 {
		query := `INSERT  INTO categories(name_category) VALUES(?)`
		_, err := db.Exec(query, categoryName)
		if err != nil {
			log.Fatal(err)
		}

	}
}
