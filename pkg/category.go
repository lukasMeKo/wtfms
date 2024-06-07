package wtfms

import (
	"github.com/meilisearch/meilisearch-go"
)

// TODO: maybe make it an interface, similiar to property. for now this is ok.
type Category struct {
	Name       string   `json:"name"`
	Properties []string `json:"properties"`
}

func GetCategory(client *meilisearch.Client, uid string) (*Category, error) {
	var category Category

	index := client.Index("categories")
	if err := index.GetDocument(uid, nil, &category); err != nil {
		return nil, err
	}

	return &category, nil
}

func (category *Category) Index() string {
	return "articles_" + category.Name
}

func RegisterCategories(client *meilisearch.Client, categories ...*Category) error {
	for _, category := range categories {
		// Could be async!
		err := ValidateCategory(client, category)
		if err != nil {
			return err
		}

		// create an index for articles of this category
		// TODO: setup index
		_, err = GetOrCreateIndex(client, category.Index(), "id")
		if err != nil {
			return err
		}

	}

	index := client.Index("categories")
	if _, err := index.AddDocuments(&categories); err != nil {
		return err
	}

	return nil
}

func ValidateCategory(client *meilisearch.Client, category *Category) error {
	// Check if properties exists
	err := PropertiesExists(client, category.Properties...)
	if err != nil {
		return err
	}

	return nil
}
