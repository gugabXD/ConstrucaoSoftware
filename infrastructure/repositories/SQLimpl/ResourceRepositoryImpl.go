package repoImpl

import (
	"database/sql"
	"sarc/core/domain"
	repositories "sarc/infrastructure/repositories/interfaces"
)

type resourceRepositoryImpl struct {
	db *sql.DB
}

func NewResourceRepository(db *sql.DB) repositories.ResourceRepository {
	return &resourceRepositoryImpl{db}
}

func (r *resourceRepositoryImpl) Create(resource *domain.Resource) error {
	return r.db.QueryRow(
		"INSERT INTO resources (description, status, characteristics, resource_type_id) VALUES ($1, $2, $3, $4) RETURNING resource_id",
		resource.Description, resource.Status, resource.Characteristics, resource.ResourceTypeID,
	).Scan(&resource.ResourceID)
}

func (r *resourceRepositoryImpl) FindAll() ([]domain.Resource, error) {
	rows, err := r.db.Query(`
        SELECT res.resource_id, res.description, res.status, res.characteristics, res.resource_type_id,
               rt.resource_type_id, rt.name
        FROM resources res
        LEFT JOIN resource_types rt ON res.resource_type_id = rt.resource_type_id
    `)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var resources []domain.Resource
	for rows.Next() {
		var res domain.Resource
		var rt domain.ResourceType
		if err := rows.Scan(&res.ResourceID, &res.Description, &res.Status, &res.Characteristics, &res.ResourceTypeID, &rt.ResourceTypeID, &rt.Name); err != nil {
			return nil, err
		}
		res.ResourceType = &rt
		resources = append(resources, res)
	}
	return resources, nil
}

func (r *resourceRepositoryImpl) FindByID(id uint) (*domain.Resource, error) {
	row := r.db.QueryRow(`
        SELECT res.resource_id, res.description, res.status, res.characteristics, res.resource_type_id,
               rt.resource_type_id, rt.name
        FROM resources res
        LEFT JOIN resource_types rt ON res.resource_type_id = rt.resource_type_id
        WHERE res.resource_id = $1
    `, id)
	var res domain.Resource
	var rt domain.ResourceType
	if err := row.Scan(&res.ResourceID, &res.Description, &res.Status, &res.Characteristics, &res.ResourceTypeID, &rt.ResourceTypeID, &rt.Name); err != nil {
		return nil, err
	}
	res.ResourceType = &rt
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
