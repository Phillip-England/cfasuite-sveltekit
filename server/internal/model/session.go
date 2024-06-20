package model

type Session struct {
	ID     int    `json:"id"`
	UserID int    `json:"user_id"`
	Token  string `json:"token"`
}

func (s Session) Table() string {
	return `
		CREATE TABLE IF NOT EXISTS session (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			user_id INTEGER,
			token TEXT,
			FOREIGN KEY (user_id) REFERENCES users(id)
		)
	`
}

func (s Session) Create() string {
	return `
		INSERT INTO session (user_id, token)
		VALUES (?, ?)
	`
}

func (s Session) Read() string {
	return `
		SELECT id, user_id, token
		FROM session
		WHERE id = ?
	`
}

func (s Session) Update() string {
	return `
		UPDATE session
		SET user_id = ?, token = ?
		WHERE id = ?
	`
}

func (s Session) Delete() string {
	return `
		DELETE FROM session
		WHERE id = ?
	`
}
