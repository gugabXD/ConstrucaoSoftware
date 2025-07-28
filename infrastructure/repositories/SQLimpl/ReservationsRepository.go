package repoImpl

import (
	"database/sql"
	"sarc/core/domain"
	repositories "sarc/infrastructure/repositories/interfaces"
)

type reservationRepositoryImpl struct {
	db *sql.DB
}

func NewReservationRepository(db *sql.DB) repositories.ReservationRepository {
	return &reservationRepositoryImpl{db}
}

func (r *reservationRepositoryImpl) Create(reservation *domain.Reservation) error {
	return r.db.QueryRow(
		"INSERT INTO reservations (lecture_id, observation) VALUES ($1, $2) RETURNING reservation_id",
		reservation.LectureID, reservation.Observation,
	).Scan(&reservation.ReservationID)
}

func (r *reservationRepositoryImpl) FindByID(id uint) (*domain.Reservation, error) {
	row := r.db.QueryRow("SELECT reservation_id, lecture_id, observation FROM reservations WHERE reservation_id = $1", id)
	var rsv domain.Reservation
	if err := row.Scan(&rsv.ReservationID, &rsv.LectureID, &rsv.Observation); err != nil {
		return nil, err
	}

	// Fetch resources for this reservation
	resRows, err := r.db.Query(`
        SELECT res.resource_id, res.description, res.status, res.characteristics, res.resource_type_id
        FROM resources res
        JOIN reservation_resources rr ON rr.resource_id = res.resource_id
        WHERE rr.reservation_id = $1
    `, id)
	if err != nil {
		return nil, err
	}
	defer resRows.Close()

	var resources []domain.Resource
	for resRows.Next() {
		var res domain.Resource
		if err := resRows.Scan(&res.ResourceID, &res.Description, &res.Status, &res.Characteristics, &res.ResourceTypeID); err != nil {
			return nil, err
		}
		resources = append(resources, res)
	}
	rsv.Resources = resources

	return &rsv, nil
}

func (r *reservationRepositoryImpl) FindAll() ([]domain.Reservation, error) {
	rows, err := r.db.Query("SELECT reservation_id, lecture_id, observation FROM reservations")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reservations []domain.Reservation
	for rows.Next() {
		var rsv domain.Reservation
		if err := rows.Scan(&rsv.ReservationID, &rsv.LectureID, &rsv.Observation); err != nil {
			return nil, err
		}

		// Fetch resources for each reservation
		resRows, err := r.db.Query(`
            SELECT res.resource_id, res.description, res.status, res.characteristics, res.resource_type_id
            FROM resources res
            JOIN reservation_resources rr ON rr.resource_id = res.resource_id
            WHERE rr.reservation_id = $1
        `, rsv.ReservationID)
		if err != nil {
			return nil, err
		}
		var resources []domain.Resource
		for resRows.Next() {
			var res domain.Resource
			if err := resRows.Scan(&res.ResourceID, &res.Description, &res.Status, &res.Characteristics, &res.ResourceTypeID); err != nil {
				resRows.Close()
				return nil, err
			}
			resources = append(resources, res)
		}
		resRows.Close()
		rsv.Resources = resources

		reservations = append(reservations, rsv)
	}
	return reservations, nil
}

func (r *reservationRepositoryImpl) Update(id uint, reservation *domain.Reservation) error {
	_, err := r.db.Exec(
		"UPDATE reservations SET lecture_id = $1, observation = $2 WHERE reservation_id = $3",
		reservation.LectureID, reservation.Observation, id,
	)
	return err
}

func (r *reservationRepositoryImpl) Delete(id uint) error {
	_, err := r.db.Exec("DELETE FROM reservations WHERE reservation_id = $1", id)
	return err
}

func (r *reservationRepositoryImpl) AddResourceToReservation(reservationID uint, resourceID uint) error {
	_, err := r.db.Exec(
		"INSERT INTO reservation_resources (reservation_id, resource_id) VALUES ($1, $2)",
		reservationID, resourceID,
	)
	return err
}
