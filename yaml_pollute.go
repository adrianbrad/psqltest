package psqltest

import (
	"database/sql"
	"github.com/romanyx/polluter"
	"io"
)

// YamlPollute is a function to insert data in a database based on a YAML.
func YamlPollute(yamlReader io.Reader, db *sql.DB) error {
	return polluter.
		New(polluter.PostgresEngine(db)).
		Pollute(yamlReader)
}
