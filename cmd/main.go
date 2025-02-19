package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/luispfcanales/confericis-backend/infraestructure/http/handlers"
	"github.com/luispfcanales/confericis-backend/infraestructure/postgres/repository"
	"github.com/luispfcanales/confericis-backend/service"
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
	http.HandleFunc("GET /api/export/svg", handlers.HandleExportSVG)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hola"))
	})

	log.Fatal(http.ListenAndServe(":3000", nil))
}
