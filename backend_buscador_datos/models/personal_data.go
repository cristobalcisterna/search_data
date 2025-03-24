package models

import (
	"context"
	"errors"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

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
		WHERE rut = $1
	`

	var record PersonalData
	err := db.QueryRow(context.Background(), query, rut).
		Scan(&record.ID, &record.Rut, &record.Name, &record.Email, &record.Phone, &record.Work)

	if err != nil {
		log.Println("Error al buscar por RUT:", err)
		return nil, errors.New("datos no encontrados")
	}

	return &record, nil
}
