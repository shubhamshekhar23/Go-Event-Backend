package db

func CreateAllTables() error {
	if err := createEventsTable(); err != nil {
		return err
	}
	if err := createUserTable(); err != nil {
		return err
	}
	if err := createRegistrationsTable(); err != nil {
		return err
	}
	return nil
}

/* Create Events Table */
func createEventsTable() error {
	query := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		date_time DATETIME NOT NULL,
		user_id INTEGER,
		FOREIGN KEY(user_id) REFERENCES users(id)
	);`
	_, err := DB.Exec(query)
	return err
}

/* Create User Table */
func createUserTable() error {
	query := `
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			email TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL
		);
	`
	_, err := DB.Exec(query)
	return err
}

/* Create User Table */
func createRegistrationsTable() error {
	query := `
		CREATE TABLE IF NOT EXISTS registrations (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			user_id INTEGER,
			event_id INTEGER,
			FOREIGN KEY(user_id) REFERENCES users(id),
			FOREIGN KEY(event_id) REFERENCES events(id)
		);
	`
	_, err := DB.Exec(query)
	return err
}
