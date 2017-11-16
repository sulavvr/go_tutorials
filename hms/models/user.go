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

func (user User) SetInfo(details map[string]string) User {
	user.name = details["name"]
	user.email = details["email"]
	user.phone = sql.NullString{String: "", Valid: false}
	user.street = sql.NullString{String: "", Valid: false}
	user.city = sql.NullString{String: "", Valid: false}
	user.state = sql.NullString{String: "", Valid: false}
	user.zip = sql.NullString{String: "", Valid: false}

	if len(details["phone"]) > 0 {
		user.phone = sql.NullString{String: details["phone"], Valid: true}
	}
	if len(details["street"]) > 0 {
		user.street = sql.NullString{String: details["street"], Valid: true}
	}
	if len(details["city"]) > 0 {
		user.city = sql.NullString{String: details["city"], Valid: true}
	}
	if len(details["state"]) > 0 {
		user.state = sql.NullString{String: details["state"], Valid: true}
	}
	if len(details["zip"]) > 0 {
		user.zip = sql.NullString{String: details["zip"], Valid: true}
	}

	return user
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

func (user User) FindByEmail(email string, db *sql.DB) (int, error) {
	query := `SELECT id FROM users WHERE email = ?`
	var uid int
	row := db.QueryRow(query, email).Scan(&uid)

	if row == sql.ErrNoRows {
		return 0, errors.New("No user found")
	}

	return uid, nil
}

func (user User) Insert(db *sql.DB) (int, error) {
	if id, err := user.FindByEmail(user.email, db); err == nil {
		return id, nil
	}

	query := `INSERT INTO users (name, email, phone, street, city, state, zip) VALUES (?, ?, ?, ?, ?, ?, ?)`
	result, err := db.Exec(query, user.name, user.email, user.phone, user.street, user.city, user.state, user.zip)

	if err != nil {
		return 0, err
	}

	last_insert, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return int(last_insert), nil
}
