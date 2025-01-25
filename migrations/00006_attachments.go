package migrations

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upAttachments, downAttachments)
}

func upAttachments(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.Exec(`--sql
		create table if not exists minichat.attachments (
			id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
			message_id uuid not null,
			filename varchar(256) not null,
			timestamp timestamptz not null DEFAULT now()
		);
	`)
	return err
}

func downAttachments(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.Exec(`--sql
		drop table if exists minichat.attachments;
	`)
	return err
}
