package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func CreateConnection() *sql.DB {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	connection_string := os.Getenv("CONNECTION_STRING")

	db, err := sql.Open("mysql", connection_string)
	if err != nil {
		log.Fatal(err)
	}

	return db
}