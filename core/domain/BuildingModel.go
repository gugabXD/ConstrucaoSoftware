package domain

type Building struct {
	BuildingID   uint   `gorm:"primaryKey" json:"buildingId"`
	BuildingName string `json:"buildingName"`
	Address      string `json:"address"`
}

type Room struct {
	RoomID       uint   `gorm:"primaryKey" json:"roomId"`
	RoomCapacity int    `json:"roomCapacity"`
	Floor        int    `json:"floor"`
	BuildingID   uint   `json:"buildingId"`
	RoomNumber   string `gorm:"uniqueIndex:idx_room_building" json:"roomNumber"`
}
