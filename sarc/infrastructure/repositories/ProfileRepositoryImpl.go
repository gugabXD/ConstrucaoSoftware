package repositories

import (
	"database/sql"
	"sarc/core/domain"
)

type profileRepositoryImpl struct {
	db *sql.DB
}

func NewProfileRepository(db *sql.DB) ProfileRepository {
	return &profileRepositoryImpl{db}
}

func (r *profileRepositoryImpl) Create(profile *domain.Profile) error {
	return r.db.QueryRow(
		"INSERT INTO profiles (role) VALUES ($1) RETURNING profile_id",
		profile.Role,
	).Scan(&profile.ID)
}

func (r *profileRepositoryImpl) FindAll() ([]domain.Profile, error) {
	rows, err := r.db.Query("SELECT profile_id, role FROM profiles")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var profiles []domain.Profile
	for rows.Next() {
		var p domain.Profile
		if err := rows.Scan(&p.ID, &p.Role); err != nil {
			return nil, err
		}
		profiles = append(profiles, p)
	}
	return profiles, nil
}

func (r *profileRepositoryImpl) FindByID(id uint) (*domain.Profile, error) {
	row := r.db.QueryRow("SELECT profile_id, role FROM profiles WHERE id = $1", id)
	var p domain.Profile
	if err := row.Scan(&p.ID, &p.Role); err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *profileRepositoryImpl) Update(id uint, profile *domain.Profile) error {
	_, err := r.db.Exec(
		"UPDATE profiles SET role = $1 WHERE id = $2",
		profile.Role, id,
	)
	return err
}

func (r *profileRepositoryImpl) Delete(id uint) error {
	_, err := r.db.Exec("DELETE FROM profiles WHERE id = $1", id)
	return err
}
