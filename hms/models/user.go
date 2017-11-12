package models

/**
 * User model contains users table row data
 */
type User struct {
	id       int
	name     string
	email    string
	password string
	phone    string
	street   string
	city     string
	state    string
	zip      string
}
