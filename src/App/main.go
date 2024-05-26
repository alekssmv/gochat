package main

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)

func main() {

	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("DATABASE_NAME"),
	)

	// Define the database connection string
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	// Ping the database to check the connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to the PostgreSQL database")

	mux := http.NewServeMux()
	// Register routes
	RegisterRoutes(mux)

	http.ListenAndServe(":8080", mux)
}
