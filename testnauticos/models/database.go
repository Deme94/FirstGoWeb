package models

import (
	// Database
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func SetupDB() {
	// Database settings
	var err error
	db, err = sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/mydatabase")
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	// Checks database connection
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}
}
