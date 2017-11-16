package models

import (
	"database/sql"
	"log"
	"strings"
)

/**
 * Room model contains rooms table row data
 */
type Room struct {
	id       int
	hotel    Hotel
	floor    int
	num      int
	roomType string
	status   bool
	smoking  bool
	rates    []Rate
}

func (room Room) GetInfo() map[string]interface{} {
	details := map[string]interface{}{
		"id":      room.id,
		"hotel":   room.hotel,
		"floor":   room.floor,
		"num":     room.num,
		"type":    strings.Title(room.roomType),
		"status":  room.status,
		"smoking": room.smoking,
		"rates":   room.rates,
	}

	return details
}

func (room Room) Rooms(hotel int, db *sql.DB) []Room {
	query := `SELECT id, hotel_id, floor, num, type, status, smoking FROM rooms WHERE hotel_id = ?`
	rows, err := db.Query(query, hotel)

	if err != nil {
		log.Fatal(err.Error())
	}
	defer rows.Close()

	var rooms []Room

	for rows.Next() {
		var h int
		rows.Scan(&room.id, &h, &room.floor, &room.num, &room.roomType, &room.status, &room.smoking)

		rate := &Rate{}
		hotel := &Hotel{}
		room.rates = rate.Rates(room.id, db)
		room.hotel = hotel.Find(h, db)
		rooms = append(rooms, room)
	}

	return rooms
}

func (room Room) Find(id int, db *sql.DB) Room {
	query := `SELECT id, hotel_id, floor, num, type, status, smoking FROM rooms WHERE id = ?`

	var h int
	row := db.QueryRow(query, id)
	row.Scan(&room.id, &h, &room.floor, &room.num, &room.roomType, &room.status, &room.smoking)

	hotel := &Hotel{}
	room.hotel = hotel.Find(h, db)

	return room
}
