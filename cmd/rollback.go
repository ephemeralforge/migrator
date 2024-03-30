/*
Copyright © 2024 Alex Crow kaitubaka@gmail.com
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/uptrace/bun/migrate"
)

// Rollback represents the migrate command
func Rollback(migrator *migrate.Migrator) *cobra.Command {
	return &cobra.Command{
		Use:   "migrate",
		Short: "migrates the last sql group",
		Run: func(cmd *cobra.Command, args []string) {
			if err := migrator.Lock(cmd.Context()); err != nil {
				fmt.Printf("cant lock the migration database\n")
			}
			defer migrator.Unlock(cmd.Context()) //nolint:errcheck

			group, err := migrator.Rollback(cmd.Context())
			if err != nil {
				fmt.Printf("cannot rollback db %s\n", err.Error())
			}

			if group.IsZero() {
				fmt.Printf("there are no new migrations to rollback (database is up to date)\n")
			}

			fmt.Printf("migrated to %s\n", group)
		},
	}
}
