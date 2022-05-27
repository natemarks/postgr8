package command

import (
	"database/sql"
	"fmt"
)

// ListDatabases Return list of all database names in the instance
func ListDatabases(conn *sql.DB) (dbNames []string, err error) {
	rows, err := conn.Query("SELECT datname FROM pg_database")

	// add databases to list
	for rows.Next() {
		var thisDB string
		if err := rows.Scan(&thisDB); err != nil {
			break
		}
		dbNames = append(dbNames, thisDB)
	}
	return dbNames, err
}

// CreateDatabase create database using an existing instance connection
func CreateDatabase(conn *sql.DB, dbName string) (err error) {
	_, err = conn.Exec(fmt.Sprintf("create database %s", dbName))
	return err
}

// DestroyDatabase destroy database using an existing instance connection
func DestroyDatabase(conn *sql.DB, dbName string) (err error) {
	_, err = conn.Exec(fmt.Sprintf("drop database %s", dbName))
	return err
}
