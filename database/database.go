package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// SetupDatabase ตั้งค่าการเชื่อมต่อฐานข้อมูล
func SetupDatabase() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// อ่านค่าจาก environment variables
	host := os.Getenv("HOST")
	port := os.Getenv("DBPORT")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DBNAME")

	// Connection string
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Open a connection
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("Unable to connect to the database:", err)
	}

	// Check the connection
	if err := db.Ping(); err != nil {
		log.Fatal("Unable to ping the database:", err)
	}

	fmt.Println("Successfully connected to the database!")
	return db
}
