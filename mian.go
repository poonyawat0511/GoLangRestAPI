package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/poonyawat0511/go-fiber/controllers"
	"github.com/poonyawat0511/go-fiber/database"

	"database/sql"
)

// ประกาศตัวแปร global สำหรับเรียกใช้ database
var db *sql.DB

func main() {

	// ตั้งค่าการเชื่อมต่อฐานข้อมูล
	db = database.SetupDatabase()
	defer db.Close()

	app := fiber.New()

	// Apply CORS middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// Register user routes
	controllers.RegisterUserRoutes(app, db)
	controllers.RegisterBookRoutes(app, db)
	controllers.RegisterUploadImage(app)

	// โหลด environment variables จากไฟล์ .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// อ่านพอร์ตจาก environment variable
	port := os.Getenv("PORT")
	if port == "" {
		port = "8888" // ใช้ค่าเริ่มต้นหากไม่ได้กำหนด PORT ใน .env
	}

	log.Printf("Starting server on port %s", port)
	log.Fatal(app.Listen(":" + port))
}
