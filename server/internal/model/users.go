package model

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u User) Table() string {
	return `
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			email TEXT,
			password BLOB
		)
	`
}

func (u User) Create() string {
	return `
		INSERT INTO users (email, password)
		VALUES (?, ?)
	`
}

func (u User) Read() string {
	return `
		SELECT id, email, password
		FROM users
		WHERE id = ?
	`
}

func (u User) Update() string {
	return `
		UPDATE users
		SET email = ?, password = ?
		WHERE id = ?
	`
}

func (u User) Delete() string {
	return `
		DELETE FROM users
		WHERE id = ?
	`
}
