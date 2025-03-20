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
	db, err := sql.Open("postgres", "postgres://postgres:hola@localhost/confericis?sslmode=disable")
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

	// Handlers
	userHandler := handlers.NewUserHandler(userCaseUse)
	roleHandler := handlers.NewRoleHandler(roleUseCase)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /users", userHandler.CreateUser)
	mux.HandleFunc("GET /api/roles", roleHandler.AllRoles)
	mux.HandleFunc("GET /api/roles/{id}", roleHandler.RoleByID)
	mux.HandleFunc("POST /api/export/svg", handlers.HandleExportSVG)
	mux.HandleFunc("POST /api/export/pdf", handlers.GeneratePDFHandler)

	handler := middleware.CorsMiddleware(middleware.LoggingMiddleware(mux))

	log.Fatal(http.ListenAndServe(":3000", handler))
}
