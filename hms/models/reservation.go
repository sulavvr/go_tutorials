package models

import (
	"time"
)

/**
 * Reservation model contains reservations table row data
 */
type Reservation struct {
	id           int
	confirmation string
	user         int
	rate         int
	checkIn      time.Time
	checkOut     time.Time
	guests       int
}
