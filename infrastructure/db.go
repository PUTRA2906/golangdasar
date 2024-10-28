package infrastructure

import (
	"database/sql"
	"log"
	"myapp/config"

	_ "github.com/go-sql-driver/mysql"
)

func NewDB() *sql.DB {
	dsn := config.GetDSN()
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	log.Println("Database connected")
	return db
}

//koneksi ke database
