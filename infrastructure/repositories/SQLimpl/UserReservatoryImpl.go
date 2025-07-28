package repoImpl

import (
	"database/sql"
	"sarc/core/domain"
	repositories "sarc/infrastructure/repositories/interfaces"
)

type userRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) repositories.UserRepository {
	return &userRepositoryImpl{db}
}

func (r *userRepositoryImpl) Create(user *domain.User) error {
	return r.db.QueryRow(
		"INSERT INTO users (email, nome, birth_date, sex, telephone, profile_id) VALUES ($1, $2, $3, $4, $5, $6) RETURNING user_id",
		user.Email, user.Nome, user.BirthDate, user.Sex, user.Telephone, user.ProfileID,
	).Scan(&user.ID)
}

func (r *userRepositoryImpl) FindAll() ([]domain.User, error) {
	rows, err := r.db.Query("SELECT user_id, email, nome, birth_date, sex, telephone, profile_id FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		var u domain.User
		if err := rows.Scan(&u.ID, &u.Email, &u.Nome, &u.BirthDate, &u.Sex, &u.Telephone, &u.ProfileID); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func (r *userRepositoryImpl) FindByID(id uint) (*domain.User, error) {
	row := r.db.QueryRow("SELECT user_id, email, nome, birth_date, sex, telephone, profile_id FROM users WHERE user_id = $1", id)
	var u domain.User
	if err := row.Scan(&u.ID, &u.Email, &u.Nome, &u.BirthDate, &u.Sex, &u.Telephone, &u.ProfileID); err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *userRepositoryImpl) Update(id uint, user *domain.User) error {
	_, err := r.db.Exec(
		"UPDATE users SET email = $1, nome = $2, birth_date = $3, sex = $4, telephone = $5, profile_id = $6 WHERE user_id = $7",
		user.Email, user.Nome, user.BirthDate, user.Sex, user.Telephone, user.ProfileID, id,
	)
	return err
}

func (r *userRepositoryImpl) Delete(id uint) error {
	_, err := r.db.Exec("DELETE FROM users WHERE user_id = $1", id)
	return err
}
