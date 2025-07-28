package repoImpl

import (
	"database/sql"
	"sarc/core/domain"
	repositories "sarc/infrastructure/repositories/interfaces"
)

type resourceTypeRepositoryImpl struct {
	db *sql.DB
}

func NewResourceTypeRepository(db *sql.DB) repositories.ResourceTypeRepository {
	return &resourceTypeRepositoryImpl{db}
}

func (r *resourceTypeRepositoryImpl) Create(resourceType *domain.ResourceType) error {
	_, err := r.db.Exec(
		"INSERT INTO resource_types (name) VALUES ($1)",
		resourceType.Name,
	)
	return err
}

func (r *resourceTypeRepositoryImpl) FindAll() ([]domain.ResourceType, error) {
	rows, err := r.db.Query("SELECT resource_type_id, name FROM resource_types")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var types []domain.ResourceType
	for rows.Next() {
		var t domain.ResourceType
		if err := rows.Scan(&t.ResourceTypeID, &t.Name); err != nil {
			return nil, err
		}
		types = append(types, t)
	}
	return types, nil
}

func (r *resourceTypeRepositoryImpl) FindByID(id uint) (*domain.ResourceType, error) {
	row := r.db.QueryRow("SELECT resource_type_id, name FROM resource_types WHERE resource_type_id = $1", id)
	var t domain.ResourceType
	if err := row.Scan(&t.ResourceTypeID, &t.Name); err != nil {
		return nil, err
	}
	return &t, nil
}

func (r *resourceTypeRepositoryImpl) Update(id uint, resourceType *domain.ResourceType) error {
	_, err := r.db.Exec(
		"UPDATE resource_types SET name = $1 WHERE resource_type_id = $2",
		resourceType.Name, id,
	)
	return err
}

func (r *resourceTypeRepositoryImpl) Delete(id uint) error {
	_, err := r.db.Exec("DELETE FROM resource_types WHERE resource_type_id = $1", id)
	return err
}
