/*
Copyright © 2024 Alex Crow kaitubaka@gmail.com
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/uptrace/bun/migrate"
)

// Init represents the init command
func Init(migrator *migrate.Migrator) *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: "create migration tables",
		Run: func(cmd *cobra.Command, args []string) {
			err := migrator.Init(cmd.Context())
			if err != nil {
				fmt.Printf("unexpected error bootstraping tables %s", err.Error())
			}
		},
	}
}
