package models

import (
	"database/sql"
	"log"
	"strings"
)

/**
 * Rate model contains rates table row data
 */
type Rate struct {
	id       int
	room     int
	rateType string
	price    int
}

func (rate Rate) GetInfo() map[string]interface{} {
	details := map[string]interface{}{
		"id":    rate.id,
		"room":  rate.room,
		"type":  strings.Title(rate.rateType),
		"price": rate.price / 100,
	}

	return details
}

func (rate Rate) Rates(room int, db *sql.DB) []Rate {
	query := `SELECT id, room_id, type, price FROM rates WHERE room_id = ?`
	rows, err := db.Query(query, room)

	if err != nil {
		log.Fatal(err.Error())
	}
	defer rows.Close()

	var rates []Rate

	for rows.Next() {
		rows.Scan(&rate.id, &rate.room, &rate.rateType, &rate.price)

		rates = append(rates, rate)
	}

	return rates
}

func (rate Rate) Find(id int, db *sql.DB) Rate {
	query := `SELECT id, room_id, type, price FROM rates WHERE id = ?`
	row := db.QueryRow(query, id)
	row.Scan(&rate.id, &rate.room, &rate.rateType, &rate.price)

	return rate
}
