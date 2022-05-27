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
// manipulate a test database named "deleteme_postgr8_test"
func TestDatabaseFunctions(t *testing.T) {
	connParams, err := internal.GetTestConnParams()
	if err != nil {
		t.Fatal("failed to get test connection parameters from AWS secrets")
	}
	conn, err := command.NewInstanceConn(connParams)
	if err != nil {
		t.Fatal("failed to get connection to test instance")
	}

	// close database at the end of the test function
	defer func(conn *sql.DB) {
		_ = conn.Close()
	}(conn)

	dbList, err := command.ListDatabases(conn)
	if err != nil {
		t.Fatal("failed to get list of databases from test instance")
	}
	if len(dbList) > 4 {
		t.Fatal("more than 4 initial databases")
	}
	// test database should  not exist yet
	if internal.StringInSlice(dbList, "deleteme_postgr8_test") {
		t.Fatal("test database should not exist yet: deleteme_postgr8_test")
	}

	// create the test database
	err = command.CreateDatabase(conn, "deleteme_postgr8_test")
	if err != nil {
		t.Fatal("failed to create test database")
	}
	// update the database list
	dbList, err = command.ListDatabases(conn)
	if err != nil {
		t.Fatal("failed to get list of databases after creating test db")
	}

	// make sure the test database exists
	if !internal.StringInSlice(dbList, "deleteme_postgr8_test") {
		t.Fatal("")
	}

	// destroy the test database
	err = command.DestroyDatabase(conn, "deleteme_postgr8_test")
	if err != nil {
		t.Fatal("failed to destroy test database")
	}

	// update the database list
	dbList, err = command.ListDatabases(conn)
	if err != nil {
		t.Fatal("failed to get list of databases after creating test db")
	}

	// make sure the test database no longer exists
	if internal.StringInSlice(dbList, "deleteme_postgr8_test") {
		t.Fatal("test database still exists after drop")
	}

}
