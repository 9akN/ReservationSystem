package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/yourname/reservation-system/internal/models"
)

type ReservationRepository struct {
	db *pgxpool.Pool
}

func NewReservationRepository(db *pgxpool.Pool) *ReservationRepository {
	return &ReservationRepository{db: db}
}

func (r *ReservationRepository) GetByID(ctx context.Context, id int32) (*models.Reservation, error) {
	query := `SELECT id, name, email, phone, date, time, guests, notes, created_at, updated_at
			  FROM reservations WHERE id = $1`

	row := r.db.QueryRow(ctx, query, id)
	res := &models.Reservation{}
	err := row.Scan(&res.ID, &res.Name, &res.Email, &res.Phone,
		&res.Date, &res.Time, &res.Guests, &res.Notes,
		&res.CreatedAt, &res.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("get reservation by id: %w", err)
	}
	return res, nil
}

func (r *ReservationRepository) List(ctx context.Context) ([]*models.Reservation, error) {
	query := `SELECT id, name, email, phone, date, time, guests, notes, created_at, updated_at
			  FROM reservations ORDER BY date, time`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("list reservations: %w", err)
	}
	defer rows.Close()

	var reservations []*models.Reservation
	for rows.Next() {
		res := &models.Reservation{}
		err := rows.Scan(&res.ID, &res.Name, &res.Email, &res.Phone,
			&res.Date, &res.Time, &res.Guests, &res.Notes,
			&res.CreatedAt, &res.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("scan reservation: %w", err)
		}
		reservations = append(reservations, res)
	}
	return reservations, nil
}

func (r *ReservationRepository) Create(ctx context.Context, req *models.CreateReservationRequest) (*models.Reservation, error) {
	date, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		return nil, fmt.Errorf("invalid date format: %w", err)
	}
	t, err := time.Parse("15:04", req.Time)
	if err != nil {
		return nil, fmt.Errorf("invalid time format: %w", err)
	}

	query := `INSERT INTO reservations (name, email, phone, date, time, guests, notes)
			  VALUES ($1, $2, $3, $4, $5, $6, $7)
			  RETURNING id, name, email, phone, date, time, guests, notes, created_at, updated_at`

	row := r.db.QueryRow(ctx, query, req.Name, req.Email, req.Phone, date, t, req.Guests, req.Notes)
	res := &models.Reservation{}
	err = row.Scan(&res.ID, &res.Name, &res.Email, &res.Phone,
		&res.Date, &res.Time, &res.Guests, &res.Notes,
		&res.CreatedAt, &res.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("create reservation: %w", err)
	}
	return res, nil
}

func (r *ReservationRepository) Update(ctx context.Context, id int32, req *models.UpdateReservationRequest) (*models.Reservation, error) {
	date, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		return nil, fmt.Errorf("invalid date format: %w", err)
	}
	t, err := time.Parse("15:04", req.Time)
	if err != nil {
		return nil, fmt.Errorf("invalid time format: %w", err)
	}

	query := `UPDATE reservations
			  SET name=$2, email=$3, phone=$4, date=$5, time=$6, guests=$7, notes=$8, updated_at=NOW()
			  WHERE id=$1
			  RETURNING id, name, email, phone, date, time, guests, notes, created_at, updated_at`

	row := r.db.QueryRow(ctx, query, id, req.Name, req.Email, req.Phone, date, t, req.Guests, req.Notes)
	res := &models.Reservation{}
	err = row.Scan(&res.ID, &res.Name, &res.Email, &res.Phone,
		&res.Date, &res.Time, &res.Guests, &res.Notes,
		&res.CreatedAt, &res.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("update reservation: %w", err)
	}
	return res, nil
}

func (r *ReservationRepository) Delete(ctx context.Context, id int32) error {
	query := `DELETE FROM reservations WHERE id = $1`
	_, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("delete reservation: %w", err)
	}
	return nil
}
