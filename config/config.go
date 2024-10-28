package config

import (
	"fmt"
	"os"
)

func GetDSN() string {
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database := os.Getenv("DB_NAME")

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, database)
}
// mengambil nilai Database dari DOT env 
