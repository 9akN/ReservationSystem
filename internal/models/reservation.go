package models

import "time"

type Reservation struct {
	ID        int32     `json:"id"         db:"id"`
	Name      string    `json:"name"       db:"name"`
	Email     string    `json:"email"      db:"email"`
	Phone     string    `json:"phone"      db:"phone"`
	Date      time.Time `json:"date"       db:"date"`
	Time      time.Time `json:"time"       db:"time"`
	Guests    int32     `json:"guests"     db:"guests"`
	Notes     *string   `json:"notes"      db:"notes"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type CreateReservationRequest struct {
	Name   string  `json:"name"`
	Email  string  `json:"email"`
	Phone  string  `json:"phone"`
	Date   string  `json:"date"` // YYYY-MM-DD
	Time   string  `json:"time"` // HH:MM
	Guests int32   `json:"guests"`
	Notes  *string `json:"notes"`
}

type UpdateReservationRequest struct {
	Name   string  `json:"name"`
	Email  string  `json:"email"`
	Phone  string  `json:"phone"`
	Date   string  `json:"date"`
	Time   string  `json:"time"`
	Guests int32   `json:"guests"`
	Notes  *string `json:"notes"`
}
