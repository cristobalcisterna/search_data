package main

import (
	"buscador_datos/db"
	"buscador_datos/routes"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	if err := db.Connect(); err != nil {
		log.Fatalf("Error conectando a la DB: %v", err)
	}
	defer db.Pool.Close()

	r := routes.SetupRouter(db.Pool)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	log.Println("Servidor corriendo en :" + port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
