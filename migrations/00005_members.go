package migrations

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upChannelMembers, downChannelMembers)
}

func upChannelMembers(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.Exec(`--sql
		create table if not exists minichat.members (
			id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
			user_id uuid not null,
			channel_id uuid not null,
			last_read_message_timestamp timestamptz not null DEFAULT now(),
			created_at timestamptz not null DEFAULT now()
		);
	`)
	return err
}

func downChannelMembers(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.Exec(`--sql
		drop table if exists minichat.members;
	`)
	return err
}
