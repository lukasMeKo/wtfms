/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"log/slog"
	"os"

	wtfms "github.com/MeKo-Tech/wtfms/pkg"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add an article",
	RunE: func(cmd *cobra.Command, args []string) error {
		categoryName, err := cmd.Flags().GetString("category")
		if err != nil {
			return err
		}

		category, err := wtfms.GetCategory(client, categoryName)
		if err != nil {
			return err
		}

		for _, articlePath := range args {
			articleData, err := os.ReadFile(articlePath)
			if err != nil {
				return err
			}

			var article wtfms.Article

			if err := json.Unmarshal(articleData, &article); err != nil {
				slog.Warn("ingoring article", "article", articlePath, "err", err)

				continue
			}

			if err := wtfms.AddArticles(client, category, &article); err != nil {
				slog.Warn("ingoring article", "article", articlePath, "err", err)

				continue
			}

			slog.Info("registered article", "category", category, "article", article)

		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.PersistentFlags().String("category", "", "the category of the article")
}
