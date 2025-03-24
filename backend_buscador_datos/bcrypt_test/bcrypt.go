package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	hash, _ := bcrypt.GenerateFromPassword([]byte("1234"), bcrypt.DefaultCost)
	fmt.Println(string(hash))
}
