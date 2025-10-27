package models

import (
	"database/sql"
	"time"
)

// Define a Snippet type to hold the data for an individual snippet. Notice how
// the fields of the struct correspon to the fields in our PostgreSQL snippets
// table?
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

// Define a SnippetModel type which wraps a sql.DB connection pool.
type SnippetModel struct {
	DB *sql.DB
}

// This will insert a new Snippet into the database.
func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	// Write the SQL statement we want to execute. I've split it over three lines
	// for readability (which is why it's surrounded by backticks instead
	// of normal double quotes).
	stmt := `INSERT INTO snippets (title, content, created, expires)
	VALUES ($1, $2, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC' + make_interval(days => $3))
	RETURNING id;`

	// Use the QueryRow() method on the embedded connection pool to execute the
	// steatement. The first parameters is the SQL statement, followed by the
	// values for the placeholder parameters: title, content, and expires in
	// that order. The method returns a *sql.Row type, which can be processed
	// with the Scan() method to load values from the returned values into the
	// supplied variable(s).
	lastInsertId := 0
	err := m.DB.QueryRow(stmt, title, content, expires).Scan(&lastInsertId)
	if err != nil {
		return 0, err
	}

	return lastInsertId, nil
}

// This will return a specific snippet based on its id.
func (m *SnippetModel) Get(id int) (Snippet, error) {
	return Snippet{}, nil
}

// This will return the 10 most recently created snippets.
func (m *SnippetModel) Latest() ([]Snippet, error) {
	return nil, nil
}
