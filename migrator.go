/*
Copyright © 2024 Alex Crow kaitubaka@gmail.com
*/
package migrator

import (
	"database/sql"

	"github.com/ephemeralforge/migrator/cmd"
	"github.com/spf13/cobra"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/extra/bundebug"
	"github.com/uptrace/bun/migrate"
	"github.com/uptrace/bun/schema"
)

type Config struct {
	DB         *sql.DB
	Migrations *migrate.Migrations
	Dialect    schema.Dialect
}

type Migrator struct {
	rootCmd *cobra.Command
}

func New(cfg *Config) *Migrator {
	bunDB := bun.NewDB(cfg.DB, cfg.Dialect)
	bunDB.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithEnabled(false),
		bundebug.FromEnv(""),
	))
	migratorClient := migrate.NewMigrator(bunDB, cfg.Migrations)
	return &Migrator{
		rootCmd: cmd.Root(
			cmd.Init(migratorClient),
			cmd.Migrate(migratorClient),
		),
	}
}

func (m *Migrator) Execute() error {
	return m.rootCmd.Execute()
}
