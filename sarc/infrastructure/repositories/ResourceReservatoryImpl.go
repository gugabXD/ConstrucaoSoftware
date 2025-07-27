package repositories

import (
	"database/sql"
	"sarc/core/domain"
)

type resourceRepositoryImpl struct {
	db *sql.DB
}

func NewResourceRepository(db *sql.DB) ResourceRepository {
	return &resourceRepositoryImpl{db}
}

func (r *resourceRepositoryImpl) Create(resource *domain.Resource) error {
	_, err := r.db.Exec(
		"INSERT INTO resources (description, status, characteristics, resource_type_id) VALUES ($1, $2, $3, $4)",
		resource.Description, resource.Status, resource.Characteristics, resource.ResourceTypeID,
	)
	return err
}

func (r *resourceRepositoryImpl) FindAll() ([]domain.Resource, error) {
	rows, err := r.db.Query("SELECT resource_id, description, status, characteristics, resource_type_id FROM resources")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var resources []domain.Resource
	for rows.Next() {
		var res domain.Resource
		if err := rows.Scan(&res.ResourceID, &res.Description, &res.Status, &res.Characteristics, &res.ResourceTypeID); err != nil {
			return nil, err
		}
		resources = append(resources, res)
	}
	return resources, nil
}

func (r *resourceRepositoryImpl) FindByID(id uint) (*domain.Resource, error) {
	row := r.db.QueryRow("SELECT resource_id, description, status, characteristics, resource_type_id FROM resources WHERE resource_id = $1", id)
	var res domain.Resource
	if err := row.Scan(&res.ResourceID, &res.Description, &res.Status, &res.Characteristics, &res.ResourceTypeID); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r *resourceRepositoryImpl) Update(id uint, resource *domain.Resource) error {
	_, err := r.db.Exec(
		"UPDATE resources SET description = $1, status = $2, characteristics = $3, resource_type_id = $4 WHERE resource_id = $5",
		resource.Description, resource.Status, resource.Characteristics, resource.ResourceTypeID, id,
	)
	return err
}

func (r *resourceRepositoryImpl) Delete(id uint) error {
	_, err := r.db.Exec("DELETE FROM resources WHERE resource_id = $1", id)
	return err
}
