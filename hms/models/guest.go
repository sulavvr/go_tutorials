package models

import (
	"time"
)

/**
 * Guest model contains guests table row data
 */
type Guest struct {
	id          int
	reservation int
	govtId      string
	idNumber    string
	checkIn     time.Time
	checkOut    time.Time
}

/*
 * AdditionalGuest model contains addl_guests table row data
 *
 */
type Additional struct {
	id          int
	reservation int
	name        string
}
