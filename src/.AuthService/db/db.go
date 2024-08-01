package db

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

var Db *gorm.DB

func Connect(maxRetries int, retryDelay time.Duration) {

	// Если не получлось загрузить .env файл, то используем переменные окружения из docker-compose.yal
	err := godotenv.Load("../../.env")
	if err != nil {
		fmt.Printf("failed to load .env file: %v", err)
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Moscow",
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_NAME"),
		os.Getenv("DATABASE_PORT"),
	)

	for i := 0; i < maxRetries; i++ {
		log.Printf("Connecting to the database (attempt %d/%d)...", i+1, maxRetries)
		// Попытка подключения к базе данных
		Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			// Успешное подключение к базе данных
			log.Println("Successfully connected to the database")
			return
		}

		// Ошибка подключения к базе данных
		log.Printf("Failed to connect to the database (attempt %d/%d): %v", i+1, maxRetries, err)
		time.Sleep(retryDelay)
	}

	// Ошибка подключения к базе данных после максимального количества попыток
	log.Fatalf("failed to connect to database after %d attempts: %v", maxRetries, err)
}
