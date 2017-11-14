package models

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
)

/**
 * User model contains users table row data
 */
type User struct {
	id       int
	name     string
	email    string
	password sql.NullString
	phone    sql.NullString
	street   sql.NullString
	city     sql.NullString
	state    sql.NullString
	zip      sql.NullString
}

func (user User) GetInfo() map[string]interface{} {
	address := "N/A"
	if user.city.Valid && user.state.Valid && user.zip.Valid {
		address = fmt.Sprintf("%s %s %s", user.city.String, user.state.String, user.zip.String)
	}

	details := map[string]interface{}{
		"id":       user.id,
		"name":     user.name,
		"email":    user.email,
		"phone":    user.phone,
		"street":   user.street,
		"location": address,
	}

	return details
}

func (user User) All(db *sql.DB) []User {
	query := `SELECT id, name, email, phone, street, city, state, zip FROM users`
	rows, err := db.Query(query)

	if err != nil {
		log.Fatal(err.Error())
	}
	defer rows.Close()

	var users []User

	for rows.Next() {
		rows.Scan(&user.id, &user.name, &user.email, &user.phone, &user.street, &user.city, &user.state, &user.zip)
		users = append(users, user)
	}

	return users
}

func (user User) Find(id int, db *sql.DB) (User, error) {
	query := `SELECT id, name, email, phone, street, city, state, zip FROM users WHERE id = ?`
	row := db.QueryRow(query, id).Scan(&user.id, &user.name, &user.email, &user.phone, &user.street, &user.city, &user.state, &user.zip)

	if row == sql.ErrNoRows {
		return user, errors.New("No user found")
	}

	return user, nil
}
