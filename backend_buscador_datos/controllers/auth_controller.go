package controllers

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"buscador_datos/models"
	"buscador_datos/utils"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AuthController struct {
	DB *pgxpool.Pool
}

func (ac *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	userModel := models.User{}
	user, err := userModel.Authenticate(ac.DB, req.Email, req.Password)
	if err != nil {
		http.Error(w, "Credenciales incorrectas", http.StatusUnauthorized)
		return
	}

	// Crear token JWT
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"rol":     user.Rol,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		http.Error(w, "Error generando token", http.StatusInternalServerError)
		return
	}

	utils.JsonResponse(w, http.StatusOK, map[string]string{
		"token": tokenString,
	})
}
