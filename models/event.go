package models

import (
	"time"

	"example.com/go-udemy-api/db"
)

type Event struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"dateTime" binding:"required"`
	UserID      int       `json:"userId"`
}

func (e Event) Save() (Event, error) {
	query := `
		INSERT INTO events (name, description, location, date_time, user_id)
		VALUES (?, ?, ?, ?, ?)
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return Event{}, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return Event{}, err
	}

	lastInsertID, _ := result.LastInsertId()
	e.ID = lastInsertID
	return e, nil
}

func GetAllEvents() ([]Event, error) {
	query := `SELECT * from events`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event

	for rows.Next() {
		var e Event
		err := rows.Scan(&e.ID, &e.Name, &e.Description, &e.Location, &e.DateTime, &e.UserID)
		if err != nil {
			return nil, err // handle scan error
		}
		events = append(events, e) // add the event to the slice
	}

	return events, nil
}

func GetEventById(id int64) (Event, error) {
	query := `SELECT * from events WHERE id = ?`
	row := db.DB.QueryRow(query, id)

	var event Event

	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		return Event{}, err // handle scan error
	}

	return event, nil
}

func UpdateEvent(e Event) error {
	query := `
		UPDATE events
		SET name = ?, description = ?, location = ?, date_time = ?, user_id = ?
		WHERE id = ?
	`
	_, err := db.DB.Exec(query, e.Name, e.Description, e.Location, e.DateTime, e.UserID, e.ID)
	return err
}

func DeleteEvent(id int64) error {
	query := `
		DELETE FROM events
		WHERE id = ?
	`
	_, err := db.DB.Exec(query, id)
	return err
}
