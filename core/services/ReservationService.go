package services

import (
	"errors"
	"sarc/core/domain"
	interfaces "sarc/core/services/interfaces"
	repositories "sarc/infrastructure/repositories/interfaces"
)

type reservationsService struct {
	repo repositories.ReservationRepository
}

func NewReservationsService(repo repositories.ReservationRepository) interfaces.ReservationsService {
	return &reservationsService{repo: repo}
}

func (s *reservationsService) CreateReservation(reservation *domain.Reservation) (*domain.Reservation, error) {
	// 1. Create the reservation itself
	if err := s.repo.Create(reservation); err != nil {
		return nil, err
	}
	// 2. Add resources to reservation_resources (many-to-many)
	for _, resource := range reservation.Resources {
		if err := s.repo.AddResourceToReservation(reservation.ReservationID, resource.ResourceID); err != nil {
			return nil, err
		}
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

func (s *reservationsService) AddResourceToReservation(reservationID uint, resourceID uint) error {
	return s.repo.AddResourceToReservation(reservationID, resourceID)
}
