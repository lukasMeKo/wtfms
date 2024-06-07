package wtfms

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/google/uuid"
	"github.com/meilisearch/meilisearch-go"
)

// TODO: make it an interface, similiar to property. for now this is ok.
type Article struct {
	Id          uuid.UUID                  `json:"id"`
	Title       string                     `json:"title"`
	Name        string                     `json:"name"`
	Description string                     `json:"description"`
	Properties  map[string]json.RawMessage `json:"properties"`
}

// TODO: batch inserts
func AddArticles(client *meilisearch.Client, category *Category, articles ...*Article) error {
	for _, article := range articles {
		article.Id = uuid.New()
		err := ValidateArticle(client, article)
		if err != nil {
			return err
		}
	}

	index := client.Index(category.Index())
	if _, err := index.AddDocuments(&articles); err != nil {
		return err
	}

	return nil
}

func ValidateArticle(client *meilisearch.Client, article *Article) error {
	for name, value := range article.Properties {
		property, err := GetProperty(client, name)
		if err != nil {
			return err
		}

		// Validate properties against schema
		validationErrs, err := schema(property).ValidateBytes(context.Background(), value)
		if err != nil {
			return err
		}

		slog.Info("validation successfull", "value", value, "property", property, "errs", validationErrs)

		if len(validationErrs) > 0 {
			return fmt.Errorf("there were validation errors: %v", validationErrs)
		}

		// Run custom validate function
		if validator, ok := property.(Validator); ok {
			err := validator.Validate(value)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
