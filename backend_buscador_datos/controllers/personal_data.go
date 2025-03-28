package controllers

import (
	"buscador_datos/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PersonalDataController struct {
	DB *pgxpool.Pool
}

// Este es el controlador que manejará la consulta por RUT
func (c *PersonalDataController) GetByRut(w http.ResponseWriter, r *http.Request) {
	// Obtener el RUT de los parámetros de la URL
	rut := mux.Vars(r)["rut"]

	// Crear una instancia del modelo de datos
	dataModel := models.PersonalData{}
	// Llamar al método del modelo para obtener los datos por el RUT
	record, err := dataModel.GetByRut(c.DB, rut)

	if err != nil {
		// Si no se encuentra el RUT, devolvemos un error
		http.Error(w, "No encontrado", http.StatusNotFound)
		return
	}

	// Si se encuentran los datos, devolvemos los datos en formato JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(record)
}

// Registrar las rutas para la consulta
func (c *PersonalDataController) RegisterRoutes(r *mux.Router) {
	personalRouter := r.PathPrefix("/personal-data").Subrouter()
	// Ruta para obtener los datos por RUT
	personalRouter.HandleFunc("/{rut}", c.GetByRut).Methods("GET")
}
