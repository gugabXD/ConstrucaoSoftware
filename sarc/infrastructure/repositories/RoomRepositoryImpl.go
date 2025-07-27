package repositories

import (
	"database/sql"
	"sarc/core/domain"
)

type roomRepositoryImpl struct {
	db *sql.DB
}

func NewRoomRepository(db *sql.DB) RoomRepository {
	return &roomRepositoryImpl{db}
}

func (r *roomRepositoryImpl) Create(room *domain.Room) error {
	_, err := r.db.Exec(
		"INSERT INTO rooms (room_number, building_id, room_capacity, floor) VALUES ($1, $2, $3, $4)",
		room.RoomNumber, room.BuildingID, room.RoomCapacity, room.Floor,
	)
	return err
}

func (r *roomRepositoryImpl) FindAll() ([]domain.Room, error) {
	rows, err := r.db.Query("SELECT room_id, room_number, building_id, room_capacity, floor FROM rooms")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rooms []domain.Room
	for rows.Next() {
		var rm domain.Room
		if err := rows.Scan(&rm.RoomID, &rm.RoomNumber, &rm.BuildingID, &rm.RoomCapacity, &rm.Floor); err != nil {
			return nil, err
		}
		rooms = append(rooms, rm)
	}
	return rooms, nil
}

func (r *roomRepositoryImpl) FindByID(id uint) (*domain.Room, error) {
	row := r.db.QueryRow("SELECT room_id, room_number, building_id, room_capacity, floor FROM rooms WHERE room_id = $1", id)
	var rm domain.Room
	if err := row.Scan(&rm.RoomID, &rm.RoomNumber, &rm.BuildingID, &rm.RoomCapacity, &rm.Floor); err != nil {
		return nil, err
	}
	return &rm, nil
}

func (r *roomRepositoryImpl) Update(id uint, room *domain.Room) error {
	_, err := r.db.Exec(
		"UPDATE rooms SET room_number = $1, building_id = $2, room_capacity = $3, floor = $4 WHERE room_id = $5",
		room.RoomNumber, room.BuildingID, room.RoomCapacity, room.Floor, id,
	)
	return err
}

func (r *roomRepositoryImpl) Delete(id uint) error {
	_, err := r.db.Exec("DELETE FROM rooms WHERE room_id = $1", id)
	return err
}
