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
	room     Room
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
		var r int
		rows.Scan(&rate.id, &r, &rate.rateType, &rate.price)

		_room := &Room{}
		rate.room = _room.Find(r, db)

		rates = append(rates, rate)
	}

	return rates
}

func (rate Rate) Find(id int, db *sql.DB) Rate {
	query := `SELECT id, room_id, type, price FROM rates WHERE id = ?`
	row := db.QueryRow(query, id)
	var r int
	row.Scan(&rate.id, &r, &rate.rateType, &rate.price)

	room := Room{}
	rate.room = room.Find(r, db)
	// log.Print(rate.room)
	return rate
}
