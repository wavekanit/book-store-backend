package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"book-store-backend/src/models"
)

var DB *gorm.DB

func ConnectDatabase() {
	// โหลดค่า Environment Variables จาก .env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// ดึง URL ฐานข้อมูลจาก .env
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL is not set in .env")
	}

	// เชื่อมต่อฐานข้อมูลด้วย GORM
	database, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// ทำ Auto Migration สำหรับ Model User
	database.AutoMigrate(&models.Users{})

	DB = database
	log.Println("Database connected successfully")
}
