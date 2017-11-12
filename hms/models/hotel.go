package models

import (
	"database/sql"
	"log"
)

/**
 * Hotel model contains hotels table row data
 */
type Hotel struct {
	id          int
	name        string
	address     string
	description sql.NullString
	rooms       int
	floors      int
	logo        sql.NullString
	price       int
	roomType    string
}

func (hotel Hotel) GetInfo() map[string]interface{} {
	details := map[string]interface{}{
		"id":          hotel.id,
		"name":        hotel.name,
		"address":     hotel.address,
		"description": hotel.description,
		"rooms":       hotel.rooms,
		"floors":      hotel.floors,
		"logo":        hotel.logo,
		"price":       hotel.price / 100,
		"type":        hotel.roomType,
	}

	return details
}

func (hotel Hotel) All(db *sql.DB) []Hotel {
	query := `SELECT t1.* FROM
            (
                SELECT hotels.id, hotels.name, hotels.address, hotels.description, hotels.floors, hotels.rooms, hotels.logo, rates.price, rates.type
                FROM hotels
                INNER JOIN rooms ON rooms.hotel_id = hotels.id
                INNER JOIN rates ON rates.room_id = rooms.id
                GROUP BY hotels.id, rates.price, rates.type
            ) t1
            LEFT OUTER JOIN
            (
                SELECT hotels.id, hotels.name, hotels.address, hotels.description, hotels.floors, hotels.rooms, hotels.logo, rates.price, rates.type
                FROM hotels
                INNER JOIN rooms ON rooms.hotel_id = hotels.id
                INNER JOIN rates ON rates.room_id = rooms.id
                GROUP BY hotels.id, rates.price, rates.type
            ) t2 ON t1.id = t2.id AND (t1.price > t2.price OR (t1.price = t2.price and t1.id < t2.id)) WHERE t2.id IS NULL`

	rows, err := db.Query(query)

	if err != nil {
		log.Fatal(err.Error())
	}
	defer rows.Close()

	var hotels []Hotel

	for rows.Next() {
		rows.Scan(&hotel.id, &hotel.name, &hotel.address, &hotel.description, &hotel.floors, &hotel.rooms, &hotel.logo, &hotel.price, &hotel.roomType)
		hotels = append(hotels, hotel)
	}

	return hotels
}

func (hotel Hotel) Find(id int, db *sql.DB) Hotel {
	query := "SELECT id, name, address, description, floors, rooms, logo FROM hotels WHERE id = ?"
	row := db.QueryRow(query, id)

	row.Scan(&hotel.id, &hotel.name, &hotel.address, &hotel.description, &hotel.floors, &hotel.rooms, &hotel.logo)

	return hotel
}
