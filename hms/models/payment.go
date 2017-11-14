package models

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	// "time"
)

/**
 * Payment model contains payments table row data
 */
type Payment struct {
	id           int
	confirmation string
	reservation  int
	amount       int
	paymentType  sql.NullString
	cardType     sql.NullString
	lastFour     sql.NullString
	paidOn       string
}

func (payment Payment) GetInfo() map[string]interface{} {
	details := map[string]interface{}{
		"id":           payment.id,
		"confirmation": payment.confirmation,
		"amount":       payment.amount / 100,
		"type":         strings.Title(payment.paymentType.String),
		"cardType":     strings.Title(payment.cardType.String),
		"lastFour":     fmt.Sprintf("***********%s", payment.lastFour.String),
		"paidOn":       payment.paidOn,
	}

	return details
}

func (payment Payment) Payment(reservation int, db *sql.DB) (Payment, error) {
	query := `SELECT id, confirmation, amount, type, card_type, card_last_four, paid_on FROM payments WHERE reservation_id = ?`
	row := db.QueryRow(query, reservation).Scan(&payment.id, &payment.confirmation, &payment.amount, &payment.paymentType, &payment.cardType, &payment.lastFour, &payment.paidOn)

	if row == sql.ErrNoRows {
		return payment, errors.New("No payment found for reservation.")
	}

	return payment, nil
}
