package service

import (
	"context"
	"fmt"

	"github.com/yourname/reservation-system/internal/models"
	"github.com/yourname/reservation-system/internal/repository"
)

type ReservationService struct {
	repo *repository.ReservationRepository
}

func NewReservationService(repo *repository.ReservationRepository) *ReservationService {
	return &ReservationService{repo: repo}
}

func (s *ReservationService) GetByID(ctx context.Context, id int32) (*models.Reservation, error) {
	if id <= 0 {
		return nil, fmt.Errorf("invalid reservation id")
	}
	return s.repo.GetByID(ctx, id)
}

func (s *ReservationService) List(ctx context.Context) ([]*models.Reservation, error) {
	return s.repo.List(ctx)
}

func (s *ReservationService) Create(ctx context.Context, req *models.CreateReservationRequest) (*models.Reservation, error) {
	if err := validateCreate(req); err != nil {
		return nil, err
	}
	return s.repo.Create(ctx, req)
}

func (s *ReservationService) Update(ctx context.Context, id int32, req *models.UpdateReservationRequest) (*models.Reservation, error) {
	if id <= 0 {
		return nil, fmt.Errorf("invalid reservation id")
	}
	if err := validateUpdate(req); err != nil {
		return nil, err
	}
	return s.repo.Update(ctx, id, req)
}

func (s *ReservationService) Delete(ctx context.Context, id int32) error {
	if id <= 0 {
		return fmt.Errorf("invalid reservation id")
	}
	return s.repo.Delete(ctx, id)
}

func validateCreate(req *models.CreateReservationRequest) error {
	if req.Name == "" {
		return fmt.Errorf("name is required")
	}
	if req.Email == "" {
		return fmt.Errorf("email is required")
	}
	if req.Phone == "" {
		return fmt.Errorf("phone is required")
	}
	if req.Date == "" {
		return fmt.Errorf("date is required")
	}
	if req.Time == "" {
		return fmt.Errorf("time is required")
	}
	if req.Guests <= 0 {
		return fmt.Errorf("guests must be greater than 0")
	}
	return nil
}

func validateUpdate(req *models.UpdateReservationRequest) error {
	if req.Name == "" {
		return fmt.Errorf("name is required")
	}
	if req.Email == "" {
		return fmt.Errorf("email is required")
	}
	if req.Phone == "" {
		return fmt.Errorf("phone is required")
	}
	if req.Date == "" {
		return fmt.Errorf("date is required")
	}
	if req.Time == "" {
		return fmt.Errorf("time is required")
	}
	if req.Guests <= 0 {
		return fmt.Errorf("guests must be greater than 0")
	}
	return nil
}
