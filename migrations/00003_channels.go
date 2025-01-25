package migrations

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upChannels, downChannels)
}

func upChannels(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.Exec(`--sql
		create table if not exists minichat.channels (
			id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
			type varchar(64) not null,
			title varchar(256) null,
      description text null,
			created_at timestamp with time zone DEFAULT now()
		);
	`)
	return err
}

func downChannels(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.Exec(`--sql
	drop table if exists minichat.channels
	`)
	return err
}
