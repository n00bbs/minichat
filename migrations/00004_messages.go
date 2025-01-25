package migrations

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upMessages, downMessages)
}

func upMessages(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.Exec(`--sql
		create table if not exists minichat.messages (
			id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
			author_id uuid not null,
			channel_id uuid not null,
			content text not null,
			timestamp timestamptz not null DEFAULT now()
		);
	`)
	return err
}

func downMessages(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.Exec(`--sql
		drop table if exists minichat.messages;
	`)
	return err
}
