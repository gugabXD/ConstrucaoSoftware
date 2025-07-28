package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"sarc/core/domain"
	"sarc/core/services"
	"sarc/infrastructure/repositories"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
	var err error
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
	fmt.Println("DSN:", dsn)
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Migrate tables with correct PK/FK names matching domain models
	_, err = DB.Exec(`
        CREATE TABLE IF NOT EXISTS profiles (
            profile_id SERIAL PRIMARY KEY,
            role TEXT NOT NULL
        );

        CREATE TABLE IF NOT EXISTS users (
            user_id SERIAL PRIMARY KEY,
            email TEXT,
            nome TEXT,
            birth_date DATE,
            sex TEXT,
            telephone TEXT,
            profile_id INTEGER REFERENCES profiles(profile_id)
        );

        CREATE TABLE IF NOT EXISTS buildings (
            building_id SERIAL PRIMARY KEY,
            building_name TEXT,
            address TEXT
        );

        CREATE TABLE IF NOT EXISTS rooms (
            room_id SERIAL PRIMARY KEY,
            room_number TEXT,
            building_id INTEGER REFERENCES buildings(building_id),
            room_capacity INTEGER,
            floor INTEGER
        );

        CREATE TABLE IF NOT EXISTS disciplines (
            discipline_id SERIAL PRIMARY KEY,
            name TEXT,
            credits INTEGER,
            program TEXT,
            bibliography TEXT[]
        );

        CREATE TABLE IF NOT EXISTS curriculums (
            curriculum_id SERIAL PRIMARY KEY,
            course_name TEXT,
            data_inicio DATE,
            data_fim DATE
        );

        CREATE TABLE IF NOT EXISTS curriculum_disciplines (
            curriculum_id INTEGER REFERENCES curriculums(curriculum_id),
            discipline_id INTEGER REFERENCES disciplines(discipline_id),
            PRIMARY KEY (curriculum_id, discipline_id)
        );

        CREATE TABLE IF NOT EXISTS classes (
            class_id SERIAL PRIMARY KEY,
            name TEXT,
            description TEXT,
            discipline_id INTEGER REFERENCES disciplines(discipline_id)
        );

        CREATE TABLE IF NOT EXISTS lectures (
            lecture_id SERIAL PRIMARY KEY,
            class_id INTEGER REFERENCES classes(class_id),
            room_id INTEGER REFERENCES rooms(room_id),
            date DATE,
            content TEXT[]
        );

        CREATE TABLE IF NOT EXISTS resource_types (
            resource_type_id SERIAL PRIMARY KEY,
            name TEXT
        );

        CREATE TABLE IF NOT EXISTS resources (
            resource_id SERIAL PRIMARY KEY,
            description TEXT,
            status TEXT,
            characteristics TEXT[],
            resource_type_id INTEGER REFERENCES resource_types(resource_type_id)
        );

        CREATE TABLE IF NOT EXISTS reservations (
            reservation_id SERIAL PRIMARY KEY,
            lecture_id INTEGER REFERENCES lectures(lecture_id),
            observation TEXT
        );

        CREATE TABLE IF NOT EXISTS reservation_resources (
            reservation_id INTEGER REFERENCES reservations(reservation_id),
            resource_id INTEGER REFERENCES resources(resource_id),
            PRIMARY KEY (reservation_id, resource_id)
        );
    `)
	if err != nil {
		log.Fatal("Failed to migrate tables:", err)
	}

	// --- Clear all tables before seeding ---
	_, err = DB.Exec(`
        TRUNCATE TABLE
            reservation_resources,
            reservations,
            resources,
            resource_types,
            lectures,
            classes,
            curriculum_disciplines,
            curriculums,
            disciplines,
            rooms,
            buildings,
            users,
            profiles
        RESTART IDENTITY CASCADE;
    `)
	if err != nil {
		log.Fatal("Failed to clear tables:", err)
	}

	// Instantiate repositories
	profileRepo := repositories.NewProfileRepository(DB)
	userRepo := repositories.NewUserRepository(DB)
	buildingRepo := repositories.NewBuildingRepository(DB)
	roomRepo := repositories.NewRoomRepository(DB)
	disciplineRepo := repositories.NewDisciplineRepository(DB)
	curriculumRepo := repositories.NewCurriculumRepository(DB)
	classRepo := repositories.NewClassRepository(DB)
	lectureRepo := repositories.NewLectureRepository(DB)
	resourceTypeRepo := repositories.NewResourceTypeRepository(DB)
	resourceRepo := repositories.NewResourceRepository(DB)
	reservationRepo := repositories.NewReservationRepository(DB)

	// Instantiate services
	profileService := services.NewProfileService(profileRepo)
	userService := services.NewUserService(userRepo)
	buildingService := services.NewBuildingService(buildingRepo)
	roomService := services.NewRoomService(roomRepo)
	disciplineService := services.NewDisciplineService(disciplineRepo)
	curriculumService := services.NewCurriculumService(curriculumRepo)
	classService := services.NewClassService(classRepo)
	lectureService := services.NewLectureService(lectureRepo)
	resourceService := services.NewResourceService(resourceRepo)
	reservationService := services.NewReservationsService(reservationRepo)

	// --- Seed data using services ---

	// Profile
	profile := &domain.Profile{Role: "admin"}
	_, err = profileService.CreateProfile(profile)
	if err != nil {
		log.Fatal("Failed to seed profile:", err)
	}

	// User
	user := &domain.User{
		Email:     "admin@example.com",
		Nome:      "Admin",
		BirthDate: "1990-01-01",
		Sex:       "M",
		Telephone: "123456789",
		ProfileID: 1,
	}
	_, err = userService.CreateUser(user)
	if err != nil {
		log.Fatal("Failed to seed user:", err)
	}

	// Building
	building := &domain.Building{
		BuildingName: "Main Building",
		Address:      "123 Main St",
	}
	_, err = buildingService.CreateBuilding(building)
	if err != nil {
		log.Fatal("Failed to seed building:", err)
	}

	// Room
	room := &domain.Room{
		RoomNumber:   "101",
		BuildingID:   1,
		RoomCapacity: 30,
		Floor:        1,
	}
	_, err = roomService.CreateRoom(room)
	if err != nil {
		log.Fatal("Failed to seed room:", err)
	}

	// Discipline
	discipline := &domain.Discipline{
		Name:         "Mathematics",
		Credits:      4,
		Program:      "Basic Math Program",
		Bibliography: []string{"Book 1", "Book 2"},
	}
	_, err = disciplineService.CreateDiscipline(discipline)
	if err != nil {
		log.Fatal("Failed to seed discipline:", err)
	}

	// Curriculum
	curriculum := &domain.Curriculum{
		CourseName: "Engineering",
		DataInicio: "2025-01-01",
		DataFim:    "2029-01-01",
		Disciplines: []domain.Discipline{
			{ID: 1}, // Add discipline by ID
		},
	}
	_, err = curriculumService.CreateCurriculum(curriculum)
	if err != nil {
		log.Fatal("Failed to seed curriculum:", err)
	}

	// Class
	class := &domain.Class{
		Name:         "Math 101",
		Description:  "Intro to Math",
		DisciplineID: 1,
	}
	_, err = classService.CreateClass(class)
	if err != nil {
		log.Fatal("Failed to seed class:", err)
	}

	// Lecture
	lecture := &domain.Lecture{
		ClassID: 1,
		RoomID:  1,
		Date:    "2025-09-01",
		Content: []string{"Introduction", "Numbers"},
	}
	_, err = lectureService.CreateLecture(lecture)
	if err != nil {
		log.Fatal("Failed to seed lecture:", err)
	}

	// ResourceType
	resourceType := &domain.ResourceType{
		Name: "Projector",
	}
	err = resourceTypeRepo.Create(resourceType)
	if err != nil {
		log.Fatal("Failed to seed resource type:", err)
	}

	// Resource
	resource := &domain.Resource{
		Description:     "Epson Projector",
		Status:          "available",
		Characteristics: []string{"HD", "HDMI"},
		ResourceTypeID:  1,
	}
	_, err = resourceService.CreateResource(resource)
	if err != nil {
		log.Fatal("Failed to seed resource:", err)
	}

	// Reservation
	reservation := &domain.Reservation{
		LectureID:   1,
		Observation: "First class reservation",
		Resources:   []domain.Resource{{ResourceID: 1}}, // Add resource by ID
	}
	_, err = reservationService.CreateReservation(reservation)
	if err != nil {
		log.Fatal("Failed to seed reservation:", err)
	}

	fmt.Println("Database connected and migrated!")
}
