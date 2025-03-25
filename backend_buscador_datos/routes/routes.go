package routes

import (
	"buscador_datos/controllers"
	"buscador_datos/middleware"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
)

func SetupRouter(db *pgxpool.Pool) *mux.Router {
	r := mux.NewRouter()

	// Controladores
	authController := controllers.AuthController{DB: db}
	dataController := controllers.StudentsPersonalDataController{DB: db}

	// Ruta pública
	r.HandleFunc("/api/login", authController.Login).Methods("POST")

	// Rutas protegidas
	api := r.PathPrefix("/api").Subrouter()
	api.Use(middleware.JWTMiddleware)

	// Registrar rutas desde controlador
	dataController.RegisterRoutes(api)

	return r
}
