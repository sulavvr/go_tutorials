package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

/**
 * Setup opens a connection to the database and pings
 * for user
 */
func Setup() {
	db, _ = sql.Open("mysql", "root@/hms")
	// defer db.Close()

	err := db.Ping()

	if err != nil {
		log.Fatal(err.Error())
	}

	log.Print("Database ready for queries.")
}

/**
 * Get returns DB instance
 */
func Get() *sql.DB {
	return db
}
