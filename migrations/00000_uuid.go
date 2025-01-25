package migrations

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upUuid, downUuid)
}

func upUuid(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.Exec(`--sql
		create extension if not exists "uuid-ossp"
	`)
	return err
}

func downUuid(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.Exec(`--sql
		drop extension if exists "uuid-ossp"
	`)
	return err
}
