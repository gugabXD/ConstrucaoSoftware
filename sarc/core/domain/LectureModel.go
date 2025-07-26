package domain

import (
	"github.com/lib/pq"
)

type Lecture struct {
	LectureID uint           `gorm:"primaryKey" json:"lectureId"`
	ClassID   uint           `json:"classId"`
	RoomID    uint           `json:"roomId"`
	Date      string         `json:"date"`
	Content   pq.StringArray `gorm:"type:text[]" json:"content" swaggertype:"array,string"`
	Presence  []User         `gorm:"many2many:lecture_presence;" json:"presence"`
}
