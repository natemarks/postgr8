package pg

import (
	"database/sql"
	"fmt"
	"github.com/natemarks/postgr8/credentials"

	_ "github.com/lib/pq"
)

//ConnectionString
func ConnectionString(creds credentials.CdkRdsAutoCredential, dbName string) string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		creds.Host,
		creds.Port,
		creds.Username,
		creds.Password,
		dbName)
}

// ValidCredentials connect to db to test credentials
func ValidCredentials(creds credentials.CdkRdsAutoCredential) (result bool, err error) {
	result = false
	// test a connection to the default maintenance database name 'postgres'
	db, err := sql.Open("postgres", ConnectionString(creds, "postgres"))
	if err != nil {
		return result, err
	}
	// close database
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)
	err = db.Ping()
	if err != nil {
		return result, err
	}

	return true, err
}
