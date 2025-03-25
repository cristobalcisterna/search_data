package models

import (
	"context"
	"encoding/json"
	"errors"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type StudentsPersonalData struct {
	Rut  *string         `json:"rut"`
	Data json.RawMessage `json:"data"`
}

func (p *StudentsPersonalData) GetByRut(db *pgxpool.Pool, rut string) (*StudentsPersonalData, error) {
	// Consulta dinámica para obtener todas las columnas como JSON
	query := `
		SELECT row_to_json(students_personal_data) 
		FROM students_personal_data 
		WHERE rut = $1
	`

	var result json.RawMessage
	// Ejecutamos la consulta y escaneamos el resultado
	err := db.QueryRow(context.Background(), query, rut).Scan(&result)

	// Verificar si hay un error en la consulta
	if err != nil {
		if err.Error() == "no rows in result set" {
			// Si no se encuentra ningún registro con ese rut, retornamos un error adecuado
			log.Println("No datos encontrados para el RUT:", rut)
			return nil, errors.New("datos no encontrados")
		}
		log.Println("Error al buscar por RUT:", err)
		return nil, err
	}

	// Retornar los datos obtenidos como JSON
	return &StudentsPersonalData{
		Rut:  &rut,
		Data: result,
	}, nil
}

/*
type PersonalData struct {
	ID    int     `json:"id"`
	Rut   *string `json:"rut"`
	Name  *string `json:"name"`
	Email *string `json:"email"`
	Phone *string `json:"phone"`
	Work  *string `json:"work"`
}

func (p *PersonalData) GetByRut(db *pgxpool.Pool, rut string) (*PersonalData, error) {
	query := `
		SELECT id, rut, name, email, phone, work
		FROM personal_data
		WHERE rut = $1`
	var record PersonalData
	err := db.QueryRow(context.Background(), query, rut).
		Scan(&record.ID, &record.Rut, &record.Name, &record.Email, &record.Phone, &record.Work)

	if err != nil {
		log.Println("Error al buscar por RUT:", err)
		return nil, errors.New("datos no encontrados")
	}

	return &record, nil
}
*/
