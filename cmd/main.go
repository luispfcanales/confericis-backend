package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/confericis-backend/infraestructure/http/handlers"
	"github.com/confericis-backend/infraestructure/postgres/repository"
	"github.com/confericis-backend/service"
	_ "github.com/lib/pq"
)

func main() {
	// Configuración de la base de datos
	db, err := sql.Open("postgres", "postgres://user:hola@localhost/dbname?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Repositorios
	userRepo := repository.NewUserRepository(db)
	roleRepo := repository.NewRoleRepository(db)

	// Servicios
	userService := service.NewUserService(userRepo, roleRepo)

	// Handlers
	userHandler := handlers.NewUserHandler(userService)

	// Rutas
	http.HandleFunc("POST /users", userHandler.CreateUser)
	http.HandleFunc("/api/export/svg", handlers.HandleExportSVG)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hola"))
	})

	log.Fatal(http.ListenAndServe(":3000", nil))
}
