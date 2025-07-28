package domain

type Class struct {
	ClassID      uint   `gorm:"primaryKey" json:"classId,omitempty" swaggerignore:"true"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	DisciplineID uint   `json:"disciplineId"`
}
