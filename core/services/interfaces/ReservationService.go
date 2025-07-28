package interfaces

import (
	"sarc/core/domain"
)

type ReservationsService interface {
	CreateReservation(reservation *domain.Reservation) (*domain.Reservation, error)
	GetReservations() ([]domain.Reservation, error)
	GetReservationByID(id uint) (*domain.Reservation, error)
	UpdateReservation(id uint, reservation *domain.Reservation) (*domain.Reservation, error)
	DeleteReservation(id uint) error
	AddResourceToReservation(reservationID uint, resourceID uint) error
}
