/*
Copyright © 2024 Alex Crow kaitubaka@gmail.com
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// Root adds all child commands to the root command and sets flags appropriately.
func Root(subCommands ...*cobra.Command) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "migrator",
		Short: "uptrace/bun migrations command line tool made in Go",
	}
	for _, subCommand := range subCommands {
		rootCmd.AddCommand(subCommand)
	}
	return rootCmd
}
