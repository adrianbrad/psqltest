package psqltest

import (
	"bytes"
	"database/sql"
	"fmt"

	"github.com/goccy/go-yaml"
	"github.com/romanyx/polluter"
)

// YamlPollute is a function to insert data in a database based on a YAML.
func YamlPollute(pollution interface{}, db *sql.DB) error {
	yamlPollution, err := yaml.Marshal(pollution)
	if err != nil {
		return fmt.Errorf("err while marshalling pollution yaml: %w", err)
	}

	return polluter.
		New(polluter.PostgresEngine(db)).
		Pollute(bytes.NewReader(yamlPollution))
}
