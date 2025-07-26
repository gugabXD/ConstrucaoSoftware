package domain

type Class struct {
	ID           uint   `gorm:"primaryKey" json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	DisciplineID uint   `json:"disciplineId"`
}
