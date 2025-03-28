package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/luispfcanales/confericis-backend/infraestructure/http/handlers"
	"github.com/luispfcanales/confericis-backend/infraestructure/postgres/repository"
	"github.com/luispfcanales/confericis-backend/middleware"
	"github.com/luispfcanales/confericis-backend/service"
)

func main() {
	// Configuración de la base de datos
	db, err := sql.Open("postgres", "postgres://postgres:hola@localhost/postgres?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	// Test the connection
	err = db.Ping()
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}
	fmt.Println("¡Hola! Successfully connected to the database")

	defer db.Close()

	// Repositorios
	userRepo := repository.NewUserRepository(db)
	roleRepo := repository.NewRoleRepository(db)

	// Servicios
	userCaseUse := service.NewUserCaseUse(userRepo, roleRepo)
	roleUseCase := service.NewRoleCaseUse(roleRepo)
	driveService := service.NewDriveService(
		"AIzaSyDdaJmn2C48NBw3O8Go50XqRlksDtnTVI0",
		"1wxxapby2lFy1GVKTbRvvRDyVha_1HHDa",
	)

	// Handlers
	userHandler := handlers.NewUserHandler(userCaseUse)
	roleHandler := handlers.NewRoleHandler(roleUseCase)
	driveHandler := handlers.NewDriveHandler(driveService)
	integrationHandler := handlers.NewIntegrationHandler(
		"https://apidatos.unamad.edu.pe/api/consulta", //oti api
		"http://200.37.144.19:6030/api/oti",           //daa api
	)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /users", userHandler.CreateUser)
	mux.HandleFunc("GET /api/roles", roleHandler.AllRoles)
	mux.HandleFunc("GET /api/roles/{id}", roleHandler.RoleByID)
	mux.HandleFunc("POST /api/export/svg", handlers.HandleExportSVG)
	mux.HandleFunc("POST /api/export/pdf", handlers.GeneratePDFHandler)
	//external apis
	mux.HandleFunc("GET /api/data/reniec/{dni}", integrationHandler.GetReniecInfo)
	mux.HandleFunc("GET /api/data/student/{code}", integrationHandler.GetStudentInfo)
	mux.HandleFunc("GET /api/data/teacher/{code}", integrationHandler.GetTeacherInfo)
	//external services
	mux.HandleFunc("GET /api/drive/files/{id}", driveHandler.ListFiles)
	mux.HandleFunc("GET /api/drive/dir", driveHandler.ListDir)

	handler := middleware.CorsMiddleware(middleware.LoggingMiddleware(mux))

	log.Fatal(http.ListenAndServe(":3000", handler))
}
