package db

import (
	"database/sql"
	"fmt"
	"log"

	"sarc/infrastructure/repositories"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
	dsn := "host=localhost user=postgres password=yourpassword dbname=mydb port=5432 sslmode=disable"
	var err error
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

	// --- Seed default data ---
	_, err = DB.Exec(`INSERT INTO profiles (role) VALUES ('admin') ON CONFLICT DO NOTHING`)
	if err != nil {
		log.Fatal("Failed to seed profile:", err)
	}

	_, err = DB.Exec(`
        INSERT INTO users (email, nome, birth_date, sex, telephone, profile_id)
        VALUES ('admin@example.com', 'Admin', '1990-01-01', 'M', '123456789', 1)
        ON CONFLICT DO NOTHING
    `)
	if err != nil {
		log.Fatal("Failed to seed user:", err)
	}

	_, err = DB.Exec(`
        INSERT INTO buildings (building_name, address)
        VALUES ('Main Building', '123 Main St')
        ON CONFLICT DO NOTHING
    `)
	if err != nil {
		log.Fatal("Failed to seed building:", err)
	}

	_, err = DB.Exec(`
        INSERT INTO rooms (room_number, building_id, room_capacity, floor)
        VALUES ('101', 1, 30, 1)
        ON CONFLICT DO NOTHING
    `)
	if err != nil {
		log.Fatal("Failed to seed room:", err)
	}

	_, err = DB.Exec(`
        INSERT INTO disciplines (name, credits, program, bibliography)
        VALUES ('Mathematics', 4, 'Basic Math Program', ARRAY['Book 1', 'Book 2'])
        ON CONFLICT DO NOTHING
    `)
	if err != nil {
		log.Fatal("Failed to seed discipline:", err)
	}

	_, err = DB.Exec(`
        INSERT INTO curriculums (course_name, data_inicio, data_fim)
        VALUES ('Engineering', '2025-01-01', '2029-01-01')
        ON CONFLICT DO NOTHING
    `)
	if err != nil {
		log.Fatal("Failed to seed curriculum:", err)
	}

	/*_, err = DB.Exec(`
	        INSERT INTO curriculum_disciplines (curriculum_id, discipline_id)
	        VALUES (1, 1)
	        ON CONFLICT DO NOTHING
	    `)
		if err != nil {
			log.Fatal("Failed to seed curriculum_disciplines:", err)
		}*/

	_, err = DB.Exec(`
        INSERT INTO classes (name, description, discipline_id)
        VALUES ('Math 101', 'Intro to Math', 1)
        ON CONFLICT DO NOTHING
    `)
	if err != nil {
		log.Fatal("Failed to seed class:", err)
	}

	_, err = DB.Exec(`
        INSERT INTO lectures (class_id, room_id, date, content)
        VALUES (1, 1, '2025-09-01', ARRAY['Introduction', 'Numbers'])
        ON CONFLICT DO NOTHING
    `)
	if err != nil {
		log.Fatal("Failed to seed lecture:", err)
	}

	_, err = DB.Exec(`
        INSERT INTO resource_types (name)
        VALUES ('Projector')
        ON CONFLICT DO NOTHING
    `)
	if err != nil {
		log.Fatal("Failed to seed resource_type:", err)
	}

	_, err = DB.Exec(`
        INSERT INTO resources (description, status, characteristics, resource_type_id)
        VALUES ('Epson Projector', 'available', ARRAY['HD', 'HDMI'], 1)
        ON CONFLICT DO NOTHING
    `)
	if err != nil {
		log.Fatal("Failed to seed resource:", err)
	}

	_, err = DB.Exec(`
        INSERT INTO reservations (lecture_id, observation)
        VALUES (1, 'First class reservation')
        ON CONFLICT DO NOTHING
    `)
	if err != nil {
		log.Fatal("Failed to seed reservation:", err)
	}

	/*_, err = DB.Exec(`
	        INSERT INTO reservation_resources (reservation_id, resource_id)
	        VALUES (1, 1)
	        ON CONFLICT DO NOTHING
	    `)
		if err != nil {
			log.Fatal("Failed to seed reservation_resources:", err)
		}*/

	// Use repository methods to connect discipline to curriculum and resource to reservation
	curriculumRepo := repositories.NewCurriculumRepository(DB)
	reservationRepo := repositories.NewReservationRepository(DB)

	// Add discipline (id=1) to curriculum (id=1)
	if err := curriculumRepo.AddDisciplineToCurriculum(1, 1); err != nil {
		log.Fatal("Failed to connect discipline to curriculum:", err)
	}

	// Add resource (id=1) to reservation (id=1)
	if err := reservationRepo.AddResourceToReservation(1, 1); err != nil {
		log.Fatal("Failed to connect resource to reservation:", err)
	}
	fmt.Println("Database connected and migrated!")
}
