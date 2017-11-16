package models

import (
	"crypto/md5"
	"database/sql"
	"fmt"
	"log"
	"time"
)

/**
 * Reservation model contains reservations table row data
 */
type Reservation struct {
	id           int
	confirmation string
	user         int
	rate         Rate
	checkIn      string
	checkOut     string
	guests       int
	hasPayment   bool
	payment      Payment
}

func (reservation Reservation) SetInfo(details map[string]interface{}) Reservation {
	data := []byte(time.Now().String())
	reservation.confirmation = fmt.Sprintf("%x", md5.Sum(data))
	reservation.user = details["user_id"].(int)
	reservation.rate = details["rate"].(Rate)
	reservation.checkIn = details["check_in"].(string)
	reservation.checkOut = details["check_out"].(string)
	reservation.guests = details["guests"].(int)

	return reservation
}

func (reservation Reservation) GetInfo() map[string]interface{} {
	details := map[string]interface{}{
		"id":           reservation.id,
		"confirmation": reservation.confirmation,
		"user":         reservation.user,
		"rate":         reservation.rate,
		"checkIn":      reservation.checkIn,
		"checkOut":     reservation.checkOut,
		"guests":       reservation.guests,
		"payment":      reservation.payment,
		"hasPayment":   reservation.hasPayment,
	}

	return details
}

func (reservation Reservation) Reservations(user int, db *sql.DB) []Reservation {
	query := "SELECT id, confirmation, rate_id, check_in, check_out, guests FROM reservations WHERE user_id = ?"
	rows, err := db.Query(query, user)

	if err != nil {
		log.Fatal(err.Error())
	}
	defer rows.Close()

	var reservations []Reservation
	var rateId int
	for rows.Next() {
		rows.Scan(&reservation.id, &reservation.confirmation, &rateId, &reservation.checkIn, &reservation.checkOut, &reservation.guests)

		rate := &Rate{}
		payment := &Payment{}
		reservation.rate = rate.Find(rateId, db)
		reservation.payment, err = payment.Payment(reservation.id, db)
		reservation.hasPayment = true

		if err != nil {
			reservation.hasPayment = false
		}
		reservations = append(reservations, reservation)
	}

	return reservations
}

func (reservation Reservation) Insert(db *sql.DB) (int, error) {
	query := `INSERT INTO reservations (confirmation, user_id, rate_id, check_in, check_out, guests)
			  VALUES (?, ?, ?, ?, ?, ?)`
	result, err := db.Exec(query, reservation.confirmation, reservation.user, reservation.rate.GetInfo()["id"], reservation.checkIn, reservation.checkOut, reservation.guests)

	if err != nil {
		return 0, err
	}

	last_insert, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return int(last_insert), nil
}
