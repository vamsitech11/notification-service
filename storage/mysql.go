// storage/mysql.go
package storage

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() {
	dsn := "root:vamsishiNnu@11@tcp(127.0.0.1:3306)/notification_service" // Replace username and password
	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Failed to connect to MySQL:", err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatal("Failed to ping MySQL:", err)
	}

	log.Println("Connected to MySQL")
}
