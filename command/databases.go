package command

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
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

// CreateDatabase
func CreateDatabase(conn *sql.DB, dbName string) (err error) {
	result, err := conn.Exec("create database " + dbName)
	fmt.Println(result.LastInsertId())
	return err
}

// DestroyDatabase
func DestroyDatabase(conn *sql.DB, dbName string) (err error) {
	result, err := conn.Exec("drop database " + dbName)
	fmt.Println(result.LastInsertId())
	return err
}
