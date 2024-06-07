/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log/slog"

	wtfms "github.com/MeKo-Tech/wtfms/pkg"
	_ "github.com/MeKo-Tech/wtfms/pkg/property"
	"github.com/spf13/cobra"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "fill some initial data",
	Run: func(cmd *cobra.Command, args []string) {
		err := wtfms.Migrate(client)
		if err != nil {
			slog.Error("faile to migrate", "err", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
