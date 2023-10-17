package db

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestConnectToDB(t *testing.T) {
	// Create a mock database connection
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock DB: %v", err)
	}
	defer db.Close()

	// Set the expectation for db.Ping to succeed
	mock.ExpectPing()

	// Define the test DSN
	testDSN := "host=localhost port=5432 user=postgres password=postgres dbname=articles sslmode=disable timezone=UTC connect_timeout=5"

	// Call ConnectToDB with the mock database connection and the actual DSN
	_, err = ConnectToDB(testDSN)

	// Check for errors
	if err != nil {
		t.Fatalf("ConnectToDB returned an error: %v", err)
	}

	// Check if the returned connection matches the mock database
	//assert.Equal(t, db, connection, "ConnectToDB did not return the expected database connection")

	// Ensure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}
func TestConnectToDBError(t *testing.T) {
	// Define a test DSN
	testDSN := "wrong databse dsn"

	// Call ConnectToDB with a test DSN known to fail
	_, err := ConnectToDB(testDSN)

	// Check for errors
	if err == nil {
		t.Fatal("ConnectToDB did not return an expected error")
	}
}
func TestOpenDBError(t *testing.T) {
	// Create a mock database connection
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock DB: %v", err)
	}
	defer db.Close()

	// Define a test DSN
	testDSN := "wrong databse dsn"

	// Call openDB with the mock database connection and test DSN
	_, err = openDB(testDSN)

	// Check for errors
	if err == nil {
		t.Fatal("openDB did not return an expected error")
	}

	// Ensure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}
