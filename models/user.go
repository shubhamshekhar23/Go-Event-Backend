package models

import (
	"example.com/go-udemy-api/db"
	"example.com/go-udemy-api/utils"
)

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u User) Save() (User, error) {
	query := `
		INSERT INTO users (email, password)
		VALUES (?, ?)
	`
	pwd, err := utils.HashPassword(u.Password)

	if err != nil {
		return User{}, err
	}

	result, err := db.DB.Exec(query, u.Email, pwd)
	if err != nil {
		return User{}, err
	}

	lastInsertID, _ := result.LastInsertId()
	u.ID = lastInsertID
	return u, nil
}

func GetAllUsers() ([]User, error) {
	query := `SELECT * from users`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User

	for rows.Next() {
		var u User
		err := rows.Scan(&u.ID, &u.Email, &u.Password)
		if err != nil {
			return nil, err // handle scan error
		}
		users = append(users, u) // add the event to the slice
	}

	return users, nil
}
