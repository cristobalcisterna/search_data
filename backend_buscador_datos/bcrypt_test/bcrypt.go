package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	hash, _ := bcrypt.GenerateFromPassword([]byte("1234"), bcrypt.DefaultCost)
	fmt.Println(string(hash))
}

//Ejecutar para generar un hash con clave "1234"
//cd ~/Escritorio/buscador_datos/backend_buscador_datos/bcrypt_test

//go run bcrypt.go
