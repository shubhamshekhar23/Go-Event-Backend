package db

func CreateAllTables() error {
	err1 := createEventsTable()
	err2 := createUserTable()

	if err1 != nil {
		return err1
	}

	if err2 != nil {
		return err2
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
