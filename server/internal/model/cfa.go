package model

type CFA struct {
	Name   string `json:"name"`
	Number int    `json:"number"`
}

func (c CFA) Table() string {
	return `
		CREATE TABLE IF NOT EXISTS cfa (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT,
			number INTEGER
		)
	`
}

func (c CFA) Create() string {
	return `
		INSERT INTO cfa (name, number)
		VALUES (?, ?)
	`
}

func (c CFA) Read() string {
	return `
		SELECT id, name, number
		FROM cfa
		WHERE id = ?
	`
}

func (c CFA) Update() string {
	return `
		UPDATE cfa
		SET name = ?, number = ?
		WHERE id = ?
	`
}

func (c CFA) Delete() string {
	return `
		DELETE FROM cfa
		WHERE id = ?
	`
}
