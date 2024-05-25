package main

import (
	"net/http"
	"time"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// connect to database
	dsn := "alex@tcp(127.0.0.1:3306)/app_db"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	mux := http.NewServeMux()
	// register routes
	RegisterRoutes(mux)
	
	http.ListenAndServe(":8080", mux)
}
