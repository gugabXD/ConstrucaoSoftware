package repositories

import (
	"database/sql"
	"sarc/core/domain"
)

type lectureRepositoryImpl struct {
	db *sql.DB
}

func NewLectureRepository(db *sql.DB) LectureRepository {
	return &lectureRepositoryImpl{db}
}

func (r *lectureRepositoryImpl) Create(lecture *domain.Lecture) error {
	return r.db.QueryRow(
		"INSERT INTO lectures (class_id, room_id, date, content) VALUES ($1, $2, $3, $4) RETURNING lecture_id",
		lecture.ClassID, lecture.RoomID, lecture.Date, lecture.Content,
	).Scan(&lecture.LectureID)
}

func (r *lectureRepositoryImpl) FindAll() ([]domain.Lecture, error) {
	rows, err := r.db.Query("SELECT lecture_id, class_id, room_id, date, content FROM lectures")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lectures []domain.Lecture
	for rows.Next() {
		var l domain.Lecture
		if err := rows.Scan(&l.LectureID, &l.ClassID, &l.RoomID, &l.Date, &l.Content); err != nil {
			return nil, err
		}
		lectures = append(lectures, l)
	}
	return lectures, nil
}

func (r *lectureRepositoryImpl) FindByID(id uint) (*domain.Lecture, error) {
	row := r.db.QueryRow("SELECT lecture_id, class_id, room_id, date, content FROM lectures WHERE lecture_id = $1", id)
	var l domain.Lecture
	if err := row.Scan(&l.LectureID, &l.ClassID, &l.RoomID, &l.Date, &l.Content); err != nil {
		return nil, err
	}
	return &l, nil
}

func (r *lectureRepositoryImpl) Update(id uint, lecture *domain.Lecture) error {
	_, err := r.db.Exec(
		"UPDATE lectures SET class_id = $1, room_id = $2, date = $3, content = $4 WHERE lecture_id = $5",
		lecture.ClassID, lecture.RoomID, lecture.Date, lecture.Content, id,
	)
	return err
}

func (r *lectureRepositoryImpl) Delete(id uint) error {
	_, err := r.db.Exec("DELETE FROM lectures WHERE lecture_id = $1", id)
	return err
}
