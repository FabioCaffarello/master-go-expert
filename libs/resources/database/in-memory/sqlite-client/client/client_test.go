package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type SQLiteClientTestSuite struct {
	suite.Suite
	client       *Client
	databasePath string
}

func TestSQLiteClientTestSuite(t *testing.T) {
	suite.Run(t, new(SQLiteClientTestSuite))
}

func (suite *SQLiteClientTestSuite) SetupTest() {
	suite.databasePath = ":memory:"
	client, err := NewClient(suite.databasePath)
	assert.NoError(suite.T(), err)
	suite.client = client
}

func (suite *SQLiteClientTestSuite) TearDownTest() {
	suite.client.db.Close()
}

func (suite *SQLiteClientTestSuite) TestNewClient() {
	client, err := NewClient(suite.databasePath)
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), client)
	assert.NotNil(suite.T(), client.db)
}

func (suite *SQLiteClientTestSuite) TestExec() {
	createTableQuery := "CREATE TABLE IF NOT EXISTS test (id INTEGER PRIMARY KEY, name TEXT)"
	err := suite.client.Exec(createTableQuery)
	assert.NoError(suite.T(), err)

	insertQuery := "INSERT INTO test (name) VALUES (?)"
	err = suite.client.Exec(insertQuery, "Alice")
	assert.NoError(suite.T(), err)
}

func (suite *SQLiteClientTestSuite) TestQueryRow() {
	createTableQuery := "CREATE TABLE IF NOT EXISTS test (id INTEGER PRIMARY KEY, name TEXT)"
	err := suite.client.Exec(createTableQuery)
	assert.NoError(suite.T(), err)

	insertQuery := "INSERT INTO test (name) VALUES (?)"
	err = suite.client.Exec(insertQuery, "Alice")
	assert.NoError(suite.T(), err)

	var name string
	query := "SELECT name FROM test WHERE id = ?"
	row := suite.client.QueryRow(query, 1)
	err = row.Scan(&name)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "Alice", name)
}

func (suite *SQLiteClientTestSuite) TestQuery() {
	createTableQuery := "CREATE TABLE IF NOT EXISTS test (id INTEGER PRIMARY KEY, name TEXT)"
	err := suite.client.Exec(createTableQuery)
	assert.NoError(suite.T(), err)

	insertQuery := "INSERT INTO test (name) VALUES (?)"
	err = suite.client.Exec(insertQuery, "Alice")
	assert.NoError(suite.T(), err)
	err = suite.client.Exec(insertQuery, "Bob")
	assert.NoError(suite.T(), err)

	query := "SELECT name FROM test"
	rows, err := suite.client.Query(query)
	assert.NoError(suite.T(), err)
	defer rows.Close()

	var names []string
	for rows.Next() {
		var name string
		err := rows.Scan(&name)
		assert.NoError(suite.T(), err)
		names = append(names, name)
	}

	assert.ElementsMatch(suite.T(), []string{"Alice", "Bob"}, names)
}
