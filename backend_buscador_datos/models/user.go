package models

import (
	"context"
	"errors"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int     `json:"id"`
	Email    *string `json:"email"`
	Password *string `json:"password"` // Solo se usa para comparar
	Rol      *string `json:"rol"`
}

// Authenticate busca el usuario por email y valida su contraseña
func (u *User) Authenticate(db *pgxpool.Pool, email string, password string) (*User, error) {
	var user User

	err := db.QueryRow(context.Background(),
		`SELECT id, email, password_hash, rol 
		 FROM users 
		 WHERE email = $1`, email).
		Scan(&user.ID, &user.Email, &user.Password, &user.Rol)

	if err != nil {
		log.Println("Error al buscar usuario:", err)
		return nil, errors.New("usuario no encontrado")
	}

	// Verificar la contraseña
	if err := bcrypt.CompareHashAndPassword([]byte(*user.Password), []byte(password)); err != nil {
		log.Println("Contraseña incorrecta:", err)
		return nil, errors.New("credenciales incorrectas")
	}

	return &user, nil
}
