package services

import (
	"errors"
	"sarc/core/domain"
	"sarc/infrastructure/repositories"
)

type ReservationsService interface {
	CreateReservation(reservation *domain.Reservation) (*domain.Reservation, error)
	GetReservations() ([]domain.Reservation, error)
	GetReservationByID(id uint) (*domain.Reservation, error)
	UpdateReservation(id uint, reservation *domain.Reservation) (*domain.Reservation, error)
	DeleteReservation(id uint) error
}

type reservationsService struct {
	repo repositories.ReservationRepository
}

func NewReservationsService(repo repositories.ReservationRepository) ReservationsService {
	return &reservationsService{repo: repo}
}

func (s *reservationsService) CreateReservation(reservation *domain.Reservation) (*domain.Reservation, error) {
	if err := s.repo.Create(reservation); err != nil {
		return nil, err
	}
	return reservation, nil
}

func (s *reservationsService) GetReservations() ([]domain.Reservation, error) {
	return s.repo.FindAll()
}

func (s *reservationsService) GetReservationByID(id uint) (*domain.Reservation, error) {
	reservation, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if reservation == nil {
		return nil, errors.New("reservation not found")
	}
	return reservation, nil
}

func (s *reservationsService) UpdateReservation(id uint, updated *domain.Reservation) (*domain.Reservation, error) {
	if err := s.repo.Update(id, updated); err != nil {
		return nil, err
	}
	return s.repo.FindByID(id)
}

func (s *reservationsService) DeleteReservation(id uint) error {
	return s.repo.Delete(id)
}
