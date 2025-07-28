package domain

type Reservation struct {
	ReservationID uint       `gorm:"primaryKey" json:"reservationId"`
	LectureID     uint       `json:"lectureId"`
	Observation   string     `json:"observation"`
	Resources     []Resource `gorm:"many2many:reservation_resources;" json:"resources"`
}
