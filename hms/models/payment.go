package models

import (
	"time"
)

/**
 * Payment model contains payments table row data
 */
type Payment struct {
	id           int
	confirmation string
	reservation  int
	amount       int
	paymentType  string
	cardType     string
	lastFour     string
	paidOn       time.Time
}
