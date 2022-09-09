package command

import (
	"database/sql"
	"fmt"
)

// InstanceConnectionParams for new postgres instance connection
// When the CDK deploys an RDS instances and automatically generates
// credentials in secretsmanager, this is the format of the JSON
type InstanceConnectionParams struct {
	Username             string `json:"username"`
	Password             string `json:"password"`
	Engine               string `json:"engine"`
	Host                 string `json:"host"`
	Port                 int    `json:"port"`
	DbInstanceIdentifier string `json:"dbInstanceIdentifier"`
}

// ConnectionString Return a connection string
// if dbName is "", return a string that connects to the instance, but not a db
// otherwise, try to connect to the instance  and the database name
func ConnectionString(connParams InstanceConnectionParams, dbName string) string {
	if dbName == "" {
		return fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable",
			connParams.Host,
			connParams.Port,
			connParams.Username,
			connParams.Password,
		)
	}

	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		connParams.Host,
		connParams.Port,
		connParams.Username,
		connParams.Password,
		dbName)
}

// ValidCredentials connect to db to test credentials
func ValidCredentials(connParams InstanceConnectionParams) (result bool, err error) {
	result = false
	// test a connection
	db, err := sql.Open("postgres", ConnectionString(connParams, ""))
	if err != nil {
		return result, err
	}
	// close database
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)
	// Use db.Ping to test the connection
	err = db.Ping()
	if err != nil {
		return result, err
	}

	return true, err
}

// NewInstanceConn Return a connection to an instance without connecting to a
// database
func NewInstanceConn(connParams InstanceConnectionParams) (conn *sql.DB, err error) {
	conn, err = sql.Open("postgres", ConnectionString(connParams, ""))
	if err != nil {
		return conn, err
	}
	return conn, err
}
