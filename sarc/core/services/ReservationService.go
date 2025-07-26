package services

import (
	"errors"

	"sarc/core/domain"

	"gorm.io/gorm"
)

// ReservationsService interface for dependency injection and testing
type ReservationsService interface {
	CreateReservation(reservation *domain.Reservation) (*domain.Reservation, error)
	GetReservations() ([]domain.Reservation, error)
	GetReservationByID(id uint) (*domain.Reservation, error)
	UpdateReservation(id uint, reservation *domain.Reservation) (*domain.Reservation, error)
	DeleteReservation(id uint) error
}

type reservationsService struct {
	db *gorm.DB
}

// NewReservationsService creates a new ReservationsService using a database connection
func NewReservationsService(db *gorm.DB) ReservationsService {
	return &reservationsService{db: db}
}

func (s *reservationsService) CreateReservation(reservation *domain.Reservation) (*domain.Reservation, error) {
	if err := s.db.Create(reservation).Error; err != nil {
		return nil, err
	}
	return reservation, nil
}

func (s *reservationsService) GetReservations() ([]domain.Reservation, error) {
	var reservations []domain.Reservation
	if err := s.db.Find(&reservations).Error; err != nil {
		return nil, err
	}
	return reservations, nil
}

func (s *reservationsService) GetReservationByID(id uint) (*domain.Reservation, error) {
	var reservation domain.Reservation
	if err := s.db.First(&reservation, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("reservation not found")
		}
		return nil, err
	}
	return &reservation, nil
}

func (s *reservationsService) UpdateReservation(id uint, updated *domain.Reservation) (*domain.Reservation, error) {
	var reservation domain.Reservation
	if err := s.db.First(&reservation, id).Error; err != nil {
		return nil, err
	}
	if err := s.db.Model(&reservation).Updates(updated).Error; err != nil {
		return nil, err
	}
	return &reservation, nil
}

func (s *reservationsService) DeleteReservation(id uint) error {
	if err := s.db.Delete(&domain.Reservation{}, id).Error; err != nil {
		return err
	}
	return nil
}
