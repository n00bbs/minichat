package migrations

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upUsers, downUsers)
}

func upUsers(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.Exec(`--sql
		create table if not exists minichat.users (
			id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
			idp_id varchar(256) not null,
			username varchar(32) not null,
			bio varchar(256),
			picture varchar(256)
		);
	`)
	return err
}

func downUsers(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.Exec(`--sql
		drop table if exists minichat.users;
	`)
	return err
}
