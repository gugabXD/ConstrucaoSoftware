package domain

type User struct {
	ID        uint   `gorm:"primaryKey" json:"id,omitempty" swaggerignore:"true"`
	Email     string `json:"email"`
	Nome      string `json:"nome"`
	BirthDate string `json:"birthDate"`
	Sex       string `json:"sex"`
	Telephone string `json:"telephone"`
	ProfileID uint   `json:"profileId"`
}

type Profile struct {
	ID   uint   `gorm:"primaryKey" json:"id,omitempty" swaggerignore:"true"`
	Role string `json:"role"`
}
