/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/meilisearch/meilisearch-go"
	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "search index",
	RunE: func(cmd *cobra.Command, args []string) error {
		query := ""
		if len(args) > 1 {
			query = args[0]
		}

		filter, _ := cmd.Flags().GetString("filter")
		index, _ := cmd.Flags().GetString("index")

		searchRes, err := client.Index(index).Search(query,
			&meilisearch.SearchRequest{
				Filter: filter,
			})
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		printJSON(searchRes)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
	searchCmd.PersistentFlags().String("index", "", "name of the index")
	searchCmd.PersistentFlags().String("filter", "", "filter")
}
