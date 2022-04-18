package psqltest

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-txdb"
)

const (
	// driver is used in order to open a new
	// transaction when opening a docker db.
	driver = "pgsqltx"

	// psqlDriver is used to register txdb.
	psqlDriver = "postgres"
)

// Register is a wrapper over txdb.Register.
// Used to register the txdb driver.
func Register(dsn string) {
	txdb.Register(driver, psqlDriver, dsn)
}

// NewTransactionTestingDB returns a new transaction DB connection.
// It acts as a test helper and also takes care of closing the DB connection.
func NewTransactionTestingDB(t *testing.T) *sql.DB {
	// mark this function as a test helper
	t.Helper()

	// open a new database connection in a SQL transaction.
	db, err := sql.Open(driver, t.Name())
	if err != nil {
		t.Fatalf("open txdb conn: %s", err)
	}

	// ping the database to ensure connection validity.
	err = db.Ping()
	if err != nil {
		t.Fatalf("ping: %s", err)
	}

	// register a cleanup function that closes the database
	// connection thus reverting the transaction.
	t.Cleanup(func() {
		_ = db.Close()
	})

	return db
}

// NewTransactionDB returns a new transaction DB connection.
func NewTransactionDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return nil, fmt.Errorf("open txdb conn: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("ping: %w", err)
	}

	return db, nil
}
