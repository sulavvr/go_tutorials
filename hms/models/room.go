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
	hotel    int
	floor    int
	num      int
	roomType string
	status   string
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
		"status":  strings.Title(room.status),
		"smoking": room.smoking,
		"rates":   room.rates,
	}

	return details
}

func (room Room) Rooms(hotel int, db *sql.DB) []Room {
	query := "SELECT id, hotel_id, floor, num, type, status, smoking FROM rooms WHERE hotel_id = ?"
	rows, err := db.Query(query, hotel)

	if err != nil {
		log.Fatal(err.Error())
	}
	defer rows.Close()

	var rooms []Room

	for rows.Next() {
		rows.Scan(&room.id, &room.hotel, &room.floor, &room.num, &room.roomType, &room.status, &room.smoking)

		r := &Rate{}
		room.rates = r.Rates(room.id, db)

		rooms = append(rooms, room)
	}

	return rooms
}
