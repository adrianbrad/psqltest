package psqltest

import (
	"bytes"
	"database/sql"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-txdb"
	"github.com/goccy/go-yaml"
	"github.com/romanyx/polluter"
)

const (
	// driver is used in order to open a new
	// transaction when opening a docker db.
	driver = "pgsqltx"

	// psqlDriver is used to register txdb.
	psqlDriver = "postgres"
)

// Pollute is a function to insert data in a database based on a YAML.
// Used only for testing.
func Pollute(pollution interface{}, db *sql.DB) error {
	yamlPollution, err := yaml.Marshal(pollution)
	if err != nil {
		return fmt.Errorf("err while marshalling pollution yaml: %w", err)
	}

	return polluter.
		New(polluter.PostgresEngine(db)).
		Pollute(bytes.NewReader(yamlPollution))
}

// Register is a wrapper over txdb.Register.
// Used to register the txdb driver.
func Register(dsn string) {
	txdb.Register(driver, psqlDriver, dsn)
}

// NewDB returns a new transaction DB connection.
func NewDB(t *testing.T) *sql.DB {
	t.Helper()

	db, err := sql.Open(driver, t.Name())
	if err != nil {
		t.Fatalf("open txdb conn: %s", err)
	}

	err = db.Ping()
	if err != nil {
		t.Fatalf("ping: %s", err)
	}

	t.Cleanup(func() {
		_ = db.Close()
	})

	return db
}
