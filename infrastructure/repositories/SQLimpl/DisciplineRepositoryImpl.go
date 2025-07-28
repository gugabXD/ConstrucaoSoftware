package repoImpl

import (
	"database/sql"
	"sarc/core/domain"
	repositories "sarc/infrastructure/repositories/interfaces"
)

type disciplineRepositoryImpl struct {
	db *sql.DB
}

func NewDisciplineRepository(db *sql.DB) repositories.DisciplineRepository {
	return &disciplineRepositoryImpl{db}
}

func (r *disciplineRepositoryImpl) Create(discipline *domain.Discipline) error {
	return r.db.QueryRow(
		"INSERT INTO disciplines (name, credits, program, bibliography) VALUES ($1, $2, $3, $4) RETURNING discipline_id",
		discipline.Name, discipline.Credits, discipline.Program, discipline.Bibliography,
	).Scan(&discipline.ID)
}

func (r *disciplineRepositoryImpl) FindAll() ([]domain.Discipline, error) {
	rows, err := r.db.Query("SELECT discipline_id, name, credits, program, bibliography FROM disciplines")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var disciplines []domain.Discipline
	for rows.Next() {
		var d domain.Discipline
		if err := rows.Scan(&d.ID, &d.Name, &d.Credits, &d.Program, &d.Bibliography); err != nil {
			return nil, err
		}
		disciplines = append(disciplines, d)
	}
	return disciplines, nil
}

func (r *disciplineRepositoryImpl) FindByID(id uint) (*domain.Discipline, error) {
	row := r.db.QueryRow("SELECT discipline_id, name, credits, program, bibliography FROM disciplines WHERE discipline_id = $1", id)
	var d domain.Discipline
	if err := row.Scan(&d.ID, &d.Name, &d.Credits, &d.Program, &d.Bibliography); err != nil {
		return nil, err
	}
	return &d, nil
}

func (r *disciplineRepositoryImpl) Update(id uint, discipline *domain.Discipline) error {
	_, err := r.db.Exec(
		"UPDATE disciplines SET name = $1, credits = $2, program = $3, bibliography = $4 WHERE discipline_id = $5",
		discipline.Name, discipline.Credits, discipline.Program, discipline.Bibliography, id,
	)
	return err
}

func (r *disciplineRepositoryImpl) Delete(id uint) error {
	_, err := r.db.Exec("DELETE FROM disciplines WHERE discipline_id = $1", id)
	return err
}
