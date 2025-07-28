package repositories

import (
	"database/sql"
	"sarc/core/domain"
)

type curriculumRepositoryImpl struct {
	db *sql.DB
}

func NewCurriculumRepository(db *sql.DB) CurriculumRepository {
	return &curriculumRepositoryImpl{db}
}

func (r *curriculumRepositoryImpl) Create(curriculum *domain.Curriculum) error {
	return r.db.QueryRow(
		"INSERT INTO curriculums (course_name, data_inicio, data_fim) VALUES ($1, $2, $3) RETURNING curriculum_id",
		curriculum.CourseName, curriculum.DataInicio, curriculum.DataFim,
	).Scan(&curriculum.ID)
}

func (r *curriculumRepositoryImpl) FindByID(id uint) (*domain.Curriculum, error) {
	row := r.db.QueryRow("SELECT curriculum_id, course_name, data_inicio, data_fim FROM curriculums WHERE curriculum_id = $1", id)
	var c domain.Curriculum
	if err := row.Scan(&c.ID, &c.CourseName, &c.DataInicio, &c.DataFim); err != nil {
		return nil, err
	}

	// Fetch disciplines for this curriculum
	discRows, err := r.db.Query(`
        SELECT d.discipline_id, d.name, d.credits, d.program, d.bibliography
        FROM disciplines d
        JOIN curriculum_disciplines cd ON cd.discipline_id = d.discipline_id
        WHERE cd.curriculum_id = $1
    `, id)
	if err != nil {
		return nil, err
	}
	defer discRows.Close()

	var disciplines []domain.Discipline
	for discRows.Next() {
		var d domain.Discipline
		if err := discRows.Scan(&d.ID, &d.Name, &d.Credits, &d.Program, &d.Bibliography); err != nil {
			return nil, err
		}
		disciplines = append(disciplines, d)
	}
	c.Disciplines = disciplines

	return &c, nil
}

func (r *curriculumRepositoryImpl) FindAll() ([]domain.Curriculum, error) {
	rows, err := r.db.Query("SELECT curriculum_id, course_name, data_inicio, data_fim FROM curriculums")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var curriculums []domain.Curriculum
	for rows.Next() {
		var c domain.Curriculum
		if err := rows.Scan(&c.ID, &c.CourseName, &c.DataInicio, &c.DataFim); err != nil {
			return nil, err
		}

		// Fetch disciplines for each curriculum
		discRows, err := r.db.Query(`
            SELECT d.discipline_id, d.name, d.credits, d.program, d.bibliography
            FROM disciplines d
            JOIN curriculum_disciplines cd ON cd.discipline_id = d.discipline_id
            WHERE cd.curriculum_id = $1
        `, c.ID)
		if err != nil {
			return nil, err
		}
		var disciplines []domain.Discipline
		for discRows.Next() {
			var d domain.Discipline
			if err := discRows.Scan(&d.ID, &d.Name, &d.Credits, &d.Program, &d.Bibliography); err != nil {
				discRows.Close()
				return nil, err
			}
			disciplines = append(disciplines, d)
		}
		discRows.Close()
		c.Disciplines = disciplines

		curriculums = append(curriculums, c)
	}
	return curriculums, nil
}

func (r *curriculumRepositoryImpl) Update(id uint, curriculum *domain.Curriculum) error {
	_, err := r.db.Exec(
		"UPDATE curriculums SET course_name = $1, data_inicio = $2, data_fim = $3 WHERE id = $4",
		curriculum.CourseName, curriculum.DataInicio, curriculum.DataFim, id,
	)
	return err
}

func (r *curriculumRepositoryImpl) Delete(id uint) error {
	_, err := r.db.Exec("DELETE FROM curriculums WHERE id = $1", id)
	return err
}

func (r *curriculumRepositoryImpl) AddDisciplineToCurriculum(curriculumID uint, disciplineID uint) error {
	_, err := r.db.Exec(
		"INSERT INTO curriculum_disciplines (curriculum_id, discipline_id) VALUES ($1, $2)",
		curriculumID, disciplineID,
	)
	return err
}
