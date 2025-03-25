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

/*
func (c *PersonalDataController) GetByRut(w http.ResponseWriter, r *http.Request) {
	rut := mux.Vars(r)["rut"]

	dataModel := models.PersonalData{}
	record, err := dataModel.GetByRut(c.DB, rut)

	if err != nil {
		http.Error(w, "No encontrado", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(record)
}
*/

func (c *PersonalDataController) GetByRut(w http.ResponseWriter, r *http.Request) {
	rut := mux.Vars(r)["rut"]

	dataModel := models.PersonalData{}
	record, err := dataModel.GetByRut(c.DB, rut)

	if err != nil {
		http.Error(w, "No encontrado", http.StatusNotFound)
		return
	}

	// Devuelve el JSON dinámico obtenido de la base de datos
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(record.Data)
}

func (c *PersonalDataController) RegisterRoutes(r *mux.Router) {
	personalRouter := r.PathPrefix("/personal-data").Subrouter()
	personalRouter.HandleFunc("/{rut}", c.GetByRut).Methods("GET")
}
