package repositories

import (
	"sarc/core/domain"

	"gorm.io/gorm"
)

type reservationRepositoryImpl struct {
	db *gorm.DB
}

func NewReservationRepository(db *gorm.DB) ReservationRepository {
	return &reservationRepositoryImpl{db}
}

func (r *reservationRepositoryImpl) Create(reservation *domain.Reservation) error {
	return r.db.Create(reservation).Error
}

func (r *reservationRepositoryImpl) FindAll() ([]domain.Reservation, error) {
	var reservations []domain.Reservation
	err := r.db.Find(&reservations).Error
	return reservations, err
}

func (r *reservationRepositoryImpl) FindByID(id uint) (*domain.Reservation, error) {
	var reservation domain.Reservation
	err := r.db.First(&reservation, id).Error
	return &reservation, err
}

func (r *reservationRepositoryImpl) Update(id uint, reservation *domain.Reservation) error {
	return r.db.Model(&domain.Reservation{}).Where("id = ?", id).Updates(reservation).Error
}

func (r *reservationRepositoryImpl) Delete(id uint) error {
	return r.db.Delete(&domain.Reservation{}, id).Error
}
