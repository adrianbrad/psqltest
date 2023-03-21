package psqltest

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-txdb"
)

const (
	// txDBDriver is used in order to open a new
	// transaction when opening a docker db.
	txDBDriver = "pgsqltx"

	// psqlDriver is used to register txdb.
	psqlDriver = "postgres"
)

// Register is a wrapper over txdb.Register.
// Used to register the txdb txDBDriver.
func Register(dsn string) {
	txdb.Register(txDBDriver, psqlDriver, dsn)
}

// RegisterWithPSQLDriver is a wrapper over txdb.Register.q
// Used to register the txdb driver with a custom psql driver.
func RegisterWithPSQLDriver(dsn, psqlDriver string) {
	txdb.Register(txDBDriver, psqlDriver, dsn)
}

// NewTransactionTestingDB returns a new transaction DB connection.
// It acts as a test helper and also takes care of closing the DB connection.
func NewTransactionTestingDB(t *testing.T) *sql.DB {
	// mark this function as a test helper
	t.Helper()

	// open a new database connection in a SQL transaction.
	db, err := sql.Open(txDBDriver, t.Name())
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
		if err := db.Close(); err != nil {
			t.Errorf("close error: %s", err)
		}
	})

	return db
}

// NewTransactionDB returns a new transaction DB connection.
func NewTransactionDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open(txDBDriver, dsn)
	if err != nil {
		return nil, fmt.Errorf("open txdb conn: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("ping: %w", err)
	}

	return db, nil
}
