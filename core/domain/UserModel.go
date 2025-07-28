package domain

type User struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Email     string `json:"email"`
	Nome      string `json:"nome"`
	BirthDate string `json:"birthDate"`
	Sex       string `json:"sex"`
	Telephone string `json:"telephone"`
	ProfileID uint   `json:"profileId"`
}

type Profile struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Role string `json:"role"`
}
