package db

import (
	"fmt"
	"log"

	"sarc/core/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=postgres password=yourpassword dbname=mydb port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto migrate your models here
	err = DB.AutoMigrate(
		&domain.User{},
		&domain.Profile{},
		&domain.Building{},
		&domain.Reservation{},
		&domain.Lecture{},
		&domain.Room{},
		&domain.Class{},
		&domain.Curriculum{},
		&domain.Discipline{},
		&domain.Resource{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// --- Seed default data ---

	// Profile
	var adminProfile domain.Profile
	DB.FirstOrCreate(&adminProfile, domain.Profile{Role: "admin"})

	// User
	var adminUser domain.User
	DB.FirstOrCreate(&adminUser, domain.User{
		Email:     "admin@example.com",
		Nome:      "Admin",
		BirthDate: "1990-01-01",
		Sex:       "M",
		Telephone: "123456789",
		ProfileID: adminProfile.ID,
	})

	// Building
	var mainBuilding domain.Building
	DB.FirstOrCreate(&mainBuilding, domain.Building{
		BuildingName: "Main Building",
		Address:      "123 Main St",
	})

	DB.Where("building_name = ?", "Main Building").First(&mainBuilding)
	// Room (connected to Building if you have a BuildingID field)
	var room domain.Room
	if err := DB.Where(domain.Room{
		RoomNumber: "101",
		BuildingID: mainBuilding.BuildingID,
	}).Assign(domain.Room{
		RoomCapacity: 30,
		Floor:        1,
	}).FirstOrCreate(&room).Error; err != nil {
		log.Fatal("Failed to seed room:", err)
	}

	// Discipline
	var discipline domain.Discipline
	DB.FirstOrCreate(&discipline, domain.Discipline{
		Name:         "Mathematics",
		Credits:      4,
		Program:      "Basic Math Program",
		Bibliography: []string{"Book 1", "Book 2"},
	})
	DB.First(&discipline, "name = ?", "Mathematics")

	// Curriculum (with discipline)
	var curriculum domain.Curriculum
	DB.FirstOrCreate(&curriculum, domain.Curriculum{
		CourseName:  "Engineering",
		DataInicio:  "2025-01-01",
		DataFim:     "2029-01-01",
		Disciplines: []domain.Discipline{discipline}, // many2many
	})
	DB.First(&curriculum, "course_name = ?", "Engineering")
	DB.Model(&curriculum).Association("Disciplines").Append(&discipline)

	// Class (connected to discipline)
	var class domain.Class
	DB.FirstOrCreate(&class, domain.Class{
		Name:         "Math 101",
		Description:  "Intro to Math",
		DisciplineID: discipline.ID,
	})

	// Lecture (connected to class and room)
	var lecture domain.Lecture
	DB.FirstOrCreate(&lecture, domain.Lecture{
		ClassID: class.ID,
		RoomID:  room.RoomID,
		Date:    "2025-09-01",
		Content: []string{"Introduction", "Numbers"},
	})

	// ResourceType
	var resourceType domain.ResourceType
	DB.FirstOrCreate(&resourceType, domain.ResourceType{
		Name: "Projector",
	})
	DB.First(&resourceType, "name = ?", "Projector")

	// Resource (connected to ResourceType)
	var resource domain.Resource
	DB.FirstOrCreate(&resource, domain.Resource{
		Description:     "Epson Projector",
		Status:          domain.ResourceStatusAvailable,
		Characteristics: []string{"HD", "HDMI"},
		ResourceTypeID:  resourceType.ResourceTypeID,
	})
	DB.First(&resource, "description = ?", "Epson Projector")

	// Reservation (connects lecture and resource)
	var reservation domain.Reservation
	DB.FirstOrCreate(&reservation, domain.Reservation{
		LectureID:   lecture.LectureID,
		Observation: "First class reservation",
		Resources:   []domain.Resource{resource}, // many2many
	})
	DB.First(&reservation, "lecture_id = ?", lecture.LectureID)
	DB.Model(&reservation).Association("Resources").Append(&resource)

	fmt.Println("Database connected and migrated!")
}
