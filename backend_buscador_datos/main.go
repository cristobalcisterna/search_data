package main

import (
	"buscador_datos/db"
	"buscador_datos/routes"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	_ = godotenv.Load()

	if err := db.Connect(); err != nil {
		log.Fatalf("Error conectando a la DB: %v", err)
	}
	defer db.Pool.Close()

	r := routes.SetupRouter(db.Pool)

	// Configuración de CORS
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"}, // Permite el frontend
		AllowedMethods: []string{"GET", "POST"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	})

	// Aplica el middleware CORS
	handler := c.Handler(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8001" // Cambié el puerto a 8001 para evitar conflicto
	}

	log.Println("Servidor corriendo en :" + port)
	log.Fatal(http.ListenAndServe(":"+port, handler)) // Usa el handler con CORS
}
