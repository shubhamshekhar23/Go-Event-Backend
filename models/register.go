package models

import (
	"errors"

	"example.com/go-udemy-api/db"
)

type Register struct {
	ID      int64 `json:"id"`
	UserID  int64 `json:"userId" binding:"required"`
	EventID int64 `json:"eventId" binding:"required"`
}

func (r Register) Save() (Register, error) {
	query := `
		INSERT INTO registrations (user_id, event_id)
		VALUES (?, ?)
	`
	result, err := db.DB.Exec(query, r.UserID, r.EventID)
	if err != nil {
		return Register{}, errors.New("registering Event Unsuccesful")
	}

	resultID, err := result.LastInsertId()
	if err != nil {
		return Register{}, err
	}
	r.ID = resultID
	return r, nil
}

func (r Register) Delete() error {
	query := `
		DELETE from registrations WHERE user_id = ? AND event_id = ?
	`
	_, err := db.DB.Exec(query, r.UserID, r.EventID)
	if err != nil {
		return errors.New("deleting Event Unsuccesful")
	}

	return nil
}
