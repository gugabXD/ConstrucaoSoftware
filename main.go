package main

import (
	"sarc/app/controllers"
	"sarc/core/services"
	_ "sarc/docs" // Importa os docs gerados
	repoimpl "sarc/infrastructure/repositories/SQLimpl"
	"sarc/pkg/db"

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	godotenv.Load()
	// Connect to the database
	db.Connect()

	// Initialize repositories
	profileRepo := repoimpl.NewProfileRepository(db.DB)
	userRepo := repoimpl.NewUserRepository(db.DB)
	buildingRepo := repoimpl.NewBuildingRepository(db.DB)
	roomRepo := repoimpl.NewRoomRepository(db.DB)
	disciplineRepo := repoimpl.NewDisciplineRepository(db.DB)
	curriculumRepo := repoimpl.NewCurriculumRepository(db.DB)
	classRepo := repoimpl.NewClassRepository(db.DB)
	lectureRepo := repoimpl.NewLectureRepository(db.DB)
	resourceRepo := repoimpl.NewResourceRepository(db.DB)
	reservationsRepo := repoimpl.NewReservationRepository(db.DB)

	// Initialize services with repositories
	buildingService := services.NewBuildingService(buildingRepo)
	roomService := services.NewRoomService(roomRepo)
	classService := services.NewClassService(classRepo)
	curriculumService := services.NewCurriculumService(curriculumRepo)
	disciplineService := services.NewDisciplineService(disciplineRepo)
	lectureService := services.NewLectureService(lectureRepo)
	profileService := services.NewProfileService(profileRepo)
	resourceService := services.NewResourceService(resourceRepo)
	userService := services.NewUserService(userRepo)
	reservationsService := services.NewReservationsService(reservationsRepo)

	// Initialize handlers
	buildingHandler := controllers.NewBuildingHandler(buildingService)
	roomHandler := controllers.NewRoomHandler(roomService)
	classHandler := controllers.NewClassHandler(classService)
	curriculumHandler := controllers.NewCurriculumHandler(curriculumService)
	disciplineHandler := controllers.NewDisciplineHandler(disciplineService)
	lectureHandler := controllers.NewLectureHandler(lectureService)
	profileHandler := controllers.NewProfileHandler(profileService)
	resourceHandler := controllers.NewResourceHandler(resourceService)
	userHandler := controllers.NewUserHandler(userService)
	reservationsHandler := controllers.NewReservationsHandler(reservationsService)

	// Setup Gin router
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Building routes
	r.POST("/buildings", buildingHandler.CreateBuilding)
	r.GET("/buildings", buildingHandler.GetBuildings)
	r.GET("/buildings/:id", buildingHandler.GetBuildingByID)
	r.PUT("/buildings/:id", buildingHandler.UpdateBuilding)
	r.DELETE("/buildings/:id", buildingHandler.DeleteBuilding)

	// Room routes (inside building or standalone)
	r.POST("/rooms", roomHandler.CreateRoom)
	r.GET("/rooms", roomHandler.GetRooms)
	r.GET("/rooms/:id", roomHandler.GetRoomByID)
	r.PUT("/rooms/:id", roomHandler.UpdateRoom)
	r.DELETE("/rooms/:id", roomHandler.DeleteRoom)

	// Class routes
	r.POST("/classes", classHandler.CreateClass)
	r.GET("/classes", classHandler.GetClasses)
	r.GET("/classes/:id", classHandler.GetClassByID)
	r.PUT("/classes/:id", classHandler.UpdateClass)
	r.DELETE("/classes/:id", classHandler.DeleteClass)

	// Curriculum routes
	r.POST("/curriculums", curriculumHandler.CreateCurriculum)
	r.GET("/curriculums", curriculumHandler.GetCurriculums)
	r.GET("/curriculums/:id", curriculumHandler.GetCurriculumByID)
	r.PUT("/curriculums/:id", curriculumHandler.UpdateCurriculum)
	r.DELETE("/curriculums/:id", curriculumHandler.DeleteCurriculum)
	r.POST("/curriculums/:id/disciplines", curriculumHandler.AddDisciplineToCurriculum)

	// Discipline routes
	r.POST("/disciplines", disciplineHandler.CreateDiscipline)
	r.GET("/disciplines", disciplineHandler.GetDisciplines)
	r.GET("/disciplines/:id", disciplineHandler.GetDisciplineByID)
	r.PUT("/disciplines/:id", disciplineHandler.UpdateDiscipline)
	r.DELETE("/disciplines/:id", disciplineHandler.DeleteDiscipline)

	// Lecture routes
	r.POST("/lectures", lectureHandler.CreateLecture)
	r.GET("/lectures", lectureHandler.GetLectures)
	r.GET("/lectures/:id", lectureHandler.GetLectureByID)
	r.PUT("/lectures/:id", lectureHandler.UpdateLecture)
	r.DELETE("/lectures/:id", lectureHandler.DeleteLecture)

	// Profile routes
	r.POST("/profiles", profileHandler.CreateProfile)
	r.GET("/profiles", profileHandler.GetProfiles)
	r.GET("/profiles/:id", profileHandler.GetProfileByID)
	r.PUT("/profiles/:id", profileHandler.UpdateProfile)
	r.DELETE("/profiles/:id", profileHandler.DeleteProfile)

	// Resource routes
	r.POST("/resources", resourceHandler.CreateResource)
	r.GET("/resources", resourceHandler.GetResources)
	r.GET("/resources/:id", resourceHandler.GetResourceByID)
	r.PUT("/resources/:id", resourceHandler.UpdateResource)
	r.DELETE("/resources/:id", resourceHandler.DeleteResource)

	// User routes
	r.POST("/users", userHandler.CreateUser)
	r.GET("/users", userHandler.GetUsers)
	r.GET("/users/:id", userHandler.GetUserByID)
	r.PUT("/users/:id", userHandler.UpdateUser)
	r.DELETE("/users/:id", userHandler.DeleteUser)

	// Reservations routes
	r.POST("/reservations", reservationsHandler.CreateReservation)
	r.GET("/reservations", reservationsHandler.GetReservations)
	r.GET("/reservations/:id", reservationsHandler.GetReservationByID)
	r.PUT("/reservations/:id", reservationsHandler.UpdateReservation)
	r.DELETE("/reservations/:id", reservationsHandler.DeleteReservation)
	r.POST("/reservations/:id/resources", reservationsHandler.AddResourceToReservation)

	// Start server
	r.Run(":8080")
}
