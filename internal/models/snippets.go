package models

import (
	"database/sql"
	"errors"
	"time"
)

type SnippetModelInterface interface {
	Insert(title string, content string, expires int) (int, error)
	Get(id int) (Snippet, error)
	Latest() ([]Snippet, error)
}

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
	// that order. The method returns a pointer to a sql.Row type, which can be
	// processed with the Scan() method to load values from the returned values
	// into the supplied variable(s).
	lastInsertId := 0
	err := m.DB.QueryRow(stmt, title, content, expires).Scan(&lastInsertId)
	if err != nil {
		return 0, err
	}

	return lastInsertId, nil
}

// This will return a specific snippet based on its id.
func (m *SnippetModel) Get(id int) (Snippet, error) {
	// Write the SQL statement we want to execute. Again, I've split it over two
	// lines for readability.
	stmt := `SELECT id, title, content, created, expires FROM snippets
	WHERE expires > NOW() AT TIME ZONE 'UTC' AND id = $1`

	// Use the QueryRow() method on the connection pool to execute our
	// SQL statement, passing in the untrusted id variable as the value for the
	// placeholder parameters. This returns a pointer to a sql.Row value which
	// holds the result from the database.
	row := m.DB.QueryRow(stmt, id)

	// Initialise a new zeroed Snippet struct.
	var s Snippet

	// Use row.Scan() to copy the values from each field in sql.Row to the
	// corresponding field in the Snippet struct. Notice that the arguments
	// to row.Scan are *pointers* to the place you want to copy the data into,
	// and the number of arguments must be exactly the same as the number of
	// columns returned by your statement.
	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err != nil {
		// If the query returns no rows, then row.Scan() will return a
		// sql.ErrNoRows error. We use the errors.Is() function to check for
		// that error specifically, and return our own ErrNoRecord error
		// instead (we'll create this in a moment).
		if errors.Is(err, sql.ErrNoRows) {
			return Snippet{}, ErrNoRecord
		} else {
			return Snippet{}, err
		}
	}

	// If everything went OK, then return the filled Snippet struct.
	return s, nil
}

// This will return the 10 most recently created snippets.
func (m *SnippetModel) Latest() ([]Snippet, error) {
	// Write the SQL statement we want to execute.
	stmt := `SELECT id, title, content, created, expires FROM snippets
	WHERE expires > now() AT TIME ZONE 'UTC' ORDER BY id DESC LIMIT 10`

	// Use the Query() method on the connection pool to execute our
	// SQL statement. This returns a sql.Rows resultset containing the result of
	// our query.
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	// We defer rows.Close() to ensure the sql.Rows resultset is
	// always properly closed before the Latest() method returns. This defer
	// statement should come *after* you check for an error from the Query()
	// method. Otherwise, if Query() returns and error, you'll get a panic
	// trying to close a nil resultset.
	defer rows.Close()

	// Initialize an empty slice to hold the Snippet structs.
	var snippets []Snippet

	// Use the rows.Next to iterate through the rows in the resultset. This
	// prepares the first (and then each subsequent) row to be acted on by the
	// rows.Scan() method. If iteration over all the rows completes then the
	// resultset autoamtically closes itself and frees up the underlying
	// database connection.
	for rows.Next() {
		// Create a new zero value Snippet struct.
		var s Snippet
		// Use rows.Scan() to copy the valued from each field in the row to the
		// new Snippet struct that we created. Again, the arguments to row.Scan()
		// must be pointers to the place you want to copy the data into, and the
		// number of arguments must be exactly the same as the number of
		// columns returned by your statement.
		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}

		// Append it to the slice of snippets.
		snippets = append(snippets, s)
	}

	// When the rows.Next() loop has finished we call rows.Err() to retrieve any
	// error that was encountered during the iteration. It's important to
	// call this - don't assume the iteration completed successfully over the
	// entire result set.
	if err = rows.Err(); err != nil {
		return nil, err
	}

	// If everything went OK then return the Snippets slice.
	return snippets, nil
}
