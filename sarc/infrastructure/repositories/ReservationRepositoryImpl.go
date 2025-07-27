package repositories

import "sarc/core/domain"

type ReservationRepository interface {
	Create(reservation *domain.Reservation) error
	FindAll() ([]domain.Reservation, error)
	FindByID(id uint) (*domain.Reservation, error)
	Update(id uint, reservation *domain.Reservation) error
	Delete(id uint) error
}
