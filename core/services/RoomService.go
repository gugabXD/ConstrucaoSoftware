package services

import (
	"errors"
	"sarc/core/domain"
	interfaces "sarc/core/services/interfaces"
	repositories "sarc/infrastructure/repositories/interfaces"
)

type roomService struct {
	repo repositories.RoomRepository
}

// NewRoomService creates a new RoomService using a repository
func NewRoomService(repo repositories.RoomRepository) interfaces.RoomService {
	return &roomService{repo: repo}
}

func (s *roomService) CreateRoom(room *domain.Room) (*domain.Room, error) {
	if err := s.repo.Create(room); err != nil {
		return nil, err
	}
	return room, nil
}

func (s *roomService) GetRooms() ([]domain.Room, error) {
	return s.repo.FindAll()
}

func (s *roomService) GetRoomByID(id uint) (*domain.Room, error) {
	room, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if room == nil {
		return nil, errors.New("room not found")
	}
	return room, nil
}

func (s *roomService) UpdateRoom(id uint, updated *domain.Room) (*domain.Room, error) {
	if err := s.repo.Update(id, updated); err != nil {
		return nil, err
	}
	return s.repo.FindByID(id)
}

func (s *roomService) DeleteRoom(id uint) error {
	return s.repo.Delete(id)
}
