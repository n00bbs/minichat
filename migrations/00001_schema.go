package migrations

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upSchema, downSchema)
}

func upSchema(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.Exec(`--sql
		create schema if not exists minichat;
	`)
	return err
}

func downSchema(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.Exec(`--sql
		drop schema if exists minichat;
	`)
	return err
}
