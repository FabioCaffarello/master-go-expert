package client

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Client struct {
	db *sql.DB
}

// NewSQLiteClient creates a new SQLite client and initializes the database.
func NewClient(dbPath string) (*Client, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}
	return &Client{
		db: db,
	}, nil
}

// Close closes the database connection.
func (c *Client) Close() error {
	return c.db.Close()
}

// Exec executes a query without returning any rows.
func (c *Client) Exec(query string, args ...interface{}) error {
	_, err := c.db.Exec(query, args...)
	return err
}

// QueryRow executes a query that is expected to return at most one row.
func (c *Client) QueryRow(query string, args ...interface{}) *sql.Row {
	return c.db.QueryRow(query, args...)
}

// Query executes a query that returns rows.
func (c *Client) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return c.db.Query(query, args...)
}

// func (c *Client) CreateCollection(name string) error {
// 	return nil
// }
