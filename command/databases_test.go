package command_test

import (
	"database/sql"
	"testing"

	"github.com/natemarks/postgr8/command"
	"github.com/natemarks/postgr8/internal"

)

// TestDatabaseFunctions create, destroy, list, etc. start with only the
// initial instance databases ("template0", "template1", "postgres",
// "rdsadmin")
func TestDatabaseFunctions(t *testing.T) {
	connParams, err := internal.GetTestConnParams()
	if err != nil{
		t.Error("failed to get test connection parameters from AWS secrets")
	}
	conn, err := command.NewInstanceConn(connParams)
	if err != nil {
		t.Error("failed to get connection to test instance")
	}


	// close database
	defer func(conn *sql.DB) {
		_ = conn.Close()
	}(conn)

	dbList, err := command.ListDatabases(conn)
	if err != nil {
		t.Error("failed to get list of databases from test instance")
	}
	if len(dbList) > 4 {
		t.Error("more than 4 initial databases")
	}
}
