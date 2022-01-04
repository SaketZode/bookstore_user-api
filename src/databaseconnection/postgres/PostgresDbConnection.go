package postgres

import (
	"bookstore_user-api/databaseconnection"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var (
	Client *sql.DB
)

type postgresDbConnection struct {
	host     string
	port     string
	username string
	password string
	schema   string
}

func newPostgresConnection() (dbclient *postgresDbConnection) {
	host := os.Getenv(databaseconnection.DB_HOST)
	port := os.Getenv(databaseconnection.DB_PORT)
	schema := os.Getenv(databaseconnection.DB_SCHEMA)
	username := os.Getenv(databaseconnection.DB_USERNAME)
	password := os.Getenv(databaseconnection.DB_PASSWORD)

	return &postgresDbConnection{
		host:     host,
		port:     port,
		schema:   schema,
		username: username,
		password: password,
	}
}

func (pgCon *postgresDbConnection) getConnectionString() (connString string) {
	connString = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", pgCon.host, pgCon.port, pgCon.username, pgCon.password, pgCon.schema)
	return connString
}

func (pgConn *postgresDbConnection) ConnectDb() {
	dataSource := pgConn.getConnectionString()
	var err error = nil
	Client, err = sql.Open("postgres", dataSource)
	if err != nil {
		panic(fmt.Sprintf("Error opening database connection!!!, %s", err))
	}
	if err = Client.Ping(); err != nil {
		panic(fmt.Sprintf("Unable to ping database!!, %s", err))
	}
}

func init() {
	dbConnection := newPostgresConnection()
	dbConnection.ConnectDb()
}
