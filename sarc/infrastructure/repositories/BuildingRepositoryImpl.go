package repositories

import (
	"database/sql"
	"sarc/core/domain"
)

type buildingRepositoryImpl struct {
	db *sql.DB
}

func NewBuildingRepository(db *sql.DB) BuildingRepository {
	return &buildingRepositoryImpl{db}
}

func (r *buildingRepositoryImpl) Create(building *domain.Building) error {
	_, err := r.db.Exec(
		"INSERT INTO buildings (building_name, address) VALUES ($1, $2)",
		building.BuildingName, building.Address,
	)
	return err
}

func (r *buildingRepositoryImpl) FindAll() ([]domain.Building, error) {
	rows, err := r.db.Query("SELECT building_id, building_name, address FROM buildings")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var buildings []domain.Building
	for rows.Next() {
		var b domain.Building
		if err := rows.Scan(&b.BuildingID, &b.BuildingName, &b.Address); err != nil {
			return nil, err
		}
		buildings = append(buildings, b)
	}
	return buildings, nil
}

func (r *buildingRepositoryImpl) FindByID(id uint) (*domain.Building, error) {
	row := r.db.QueryRow("SELECT building_id, building_name, address FROM buildings WHERE id = $1", id)
	var b domain.Building
	if err := row.Scan(&b.BuildingID, &b.BuildingName, &b.Address); err != nil {
		return nil, err
	}
	return &b, nil
}

func (r *buildingRepositoryImpl) Update(id uint, building *domain.Building) error {
	_, err := r.db.Exec(
		"UPDATE buildings SET building_name = $1, address = $2 WHERE id = $3",
		building.BuildingName, building.Address, id,
	)
	return err
}

func (r *buildingRepositoryImpl) Delete(id uint) error {
	_, err := r.db.Exec("DELETE FROM buildings WHERE id = $1", id)
	return err
}
