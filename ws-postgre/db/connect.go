package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func ConnectionToDatabase() *sql.DB {
	dns := "user=postgres password=anangs port=5432 host=localhost dbname=ws_chat_train sslmode=disable TimeZone=Asia/Jakarta"
	db, err := sql.Open("postgres", dns)
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	return db
}
