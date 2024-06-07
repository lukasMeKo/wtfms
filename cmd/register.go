package cmd

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"

	wtfms "github.com/MeKo-Tech/wtfms/pkg"
	"github.com/spf13/cobra"
)

type Type string

var (
	TypeProperty Type = "property"
	TypeCategory Type = "category"
)

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "register additional schemas",
	RunE: func(cmd *cobra.Command, args []string) error {
		typeName, err := cmd.Flags().GetString("type")
		if err != nil {
			return err
		}

		for _, schemaPath := range args {
			schema, err := os.ReadFile(schemaPath)
			if err != nil {
				return err
			}

			switch Type(typeName) {
			case TypeProperty:
				var property wtfms.PropertyType

				if err := json.Unmarshal(schema, &property); err != nil {
					slog.Warn("ingoring schema", "schema", schemaPath, "err", err)

					continue
				}

				if err := wtfms.RegisterProperties(client, property); err != nil {
					slog.Warn("ingoring schema", "schema", schemaPath, "err", err)

					continue
				}

				slog.Info("registered property", "category", property)

			case TypeCategory:
				var category wtfms.Category

				if err := json.Unmarshal(schema, &category); err != nil {
					slog.Warn("ingoring schema", "schema", schemaPath, "err", err)

					continue
				}

				if err := wtfms.RegisterCategories(client, &category); err != nil {
					slog.Warn("ingoring schema", "schema", schemaPath, "err", err)

					continue
				}

				slog.Info("registered category", "category", category)

			default:
				return fmt.Errorf("unknown type: '%s'", typeName)
			}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)
	registerCmd.PersistentFlags().String("type", "", "type of the schema")
}
