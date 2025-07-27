package repositories

import (
	"database/sql"
	"sarc/core/domain"
)

type classRepositoryImpl struct {
	db *sql.DB
}

func NewClassRepository(db *sql.DB) ClassRepository {
	return &classRepositoryImpl{db}
}

func (r *classRepositoryImpl) Create(class *domain.Class) error {
	_, err := r.db.Exec(
		"INSERT INTO classes (name, description, discipline_id) VALUES ($1, $2, $3)",
		class.Name, class.Description, class.DisciplineID,
	)
	return err
}

func (r *classRepositoryImpl) FindAll() ([]domain.Class, error) {
	rows, err := r.db.Query("SELECT class_id, name, description, discipline_id FROM classes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var classes []domain.Class
	for rows.Next() {
		var c domain.Class
		if err := rows.Scan(&c.ID, &c.Name, &c.Description, &c.DisciplineID); err != nil {
			return nil, err
		}
		classes = append(classes, c)
	}
	return classes, nil
}

func (r *classRepositoryImpl) FindByID(id uint) (*domain.Class, error) {
	row := r.db.QueryRow("SELECT class_id, name, description, discipline_id FROM classes WHERE id = $1", id)
	var c domain.Class
	if err := row.Scan(&c.ID, &c.Name, &c.Description, &c.DisciplineID); err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *classRepositoryImpl) Update(id uint, class *domain.Class) error {
	_, err := r.db.Exec(
		"UPDATE classes SET name = $1, description = $2, discipline_id = $3 WHERE id = $4",
		class.Name, class.Description, class.DisciplineID, id,
	)
	return err
}

func (r *classRepositoryImpl) Delete(id uint) error {
	_, err := r.db.Exec("DELETE FROM classes WHERE id = $1", id)
	return err
}
