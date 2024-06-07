package cmd

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"

	"github.com/meilisearch/meilisearch-go"
	"github.com/spf13/cobra"
)

var client = meilisearch.NewClient(meilisearch.ClientConfig{
	Host:   "http://localhost:7700",
	APIKey: "masterKey",
})

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "wtfms",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		AddSource: false,
		Level:     slog.LevelDebug,
	})))
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.polymeko-wtfms.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func printJSON(v any) {
	pretty, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		fmt.Printf("%+v", v)
	}
	fmt.Println(string(pretty))
}
