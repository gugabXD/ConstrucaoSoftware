package domain

import (
	"github.com/lib/pq"
)

type Discipline struct {
	ID           uint           `gorm:"primaryKey" json:"id,omitempty" swaggerignore:"true"`
	Name         string         `json:"name"`
	Credits      int            `json:"credits"`
	Program      string         `json:"program"`
	Bibliography pq.StringArray `gorm:"type:text[]" json:"bibliography" swaggertype:"array,string"`
}

type Curriculum struct {
	ID          uint         `gorm:"primaryKey" json:"id,omitempty" swaggerignore:"true"`
	CourseName  string       `json:"courseName"`
	DataInicio  string       `json:"dataInicio"`
	DataFim     string       `json:"dataFim"`
	Disciplines []Discipline `gorm:"many2many:curriculum_disciplines;" json:"disciplines"`
}
