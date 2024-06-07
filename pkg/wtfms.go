package wtfms

import (
	"log/slog"

	"github.com/invopop/jsonschema"
	"github.com/meilisearch/meilisearch-go"
)

var (
	// properties defined in go code.
	properties map[string]Property
)

func RegisterProperty(property Property) {
	if properties == nil {
		properties = make(map[string]Property)
	}

	if existingProperty, found := properties[property.Name()]; found {
		slog.Warn("overriding existing property",
			"name", property.Name,
			"property", existingProperty,
		)
	}

	properties[property.Name()] = property
}

func Migrate(client *meilisearch.Client) error {
	/* TODO:
	- make only name searchable
	*/
	slog.Info("running migrations", "properties", properties, "len", len(properties))

	_, err := GetOrCreateIndex(client, "categories", "name")
	if err != nil {
		return err
	}

	index, err := GetOrCreateIndex(client, "properties", "name")
	if err != nil {
		return err
	}

	// TODO: use diffing
	if _, err = index.DeleteAllDocuments(); err != nil {
		return err
	}

	var documents []map[string]any

	// Add registered properties
	for name, property := range properties {
		schema := jsonschema.Reflect(property)
		documents = append(documents, map[string]any{
			"name":   name,
			"schema": schema,
		})
		slog.Info("creating property", "name", name, "schema", schema)
	}

	_, err = index.AddDocuments(documents, "name")
	if err != nil {
		return err
	}

	slog.Info("deleted old documents")

	return nil
}

func GetOrCreateIndex(client *meilisearch.Client, uid, primaryKey string) (*meilisearch.Index, error) {
	// Create Index
	createIndexTask, err := client.CreateIndex(&meilisearch.IndexConfig{
		Uid:        uid,
		PrimaryKey: primaryKey,
	})
	if err != nil {
		return nil, err
	}

	if _, err = client.WaitForTask(createIndexTask.TaskUID); err != nil {
		return nil, err
	}

	// Get the created index
	index, err := client.GetIndex(uid)
	if err != nil {
		return nil, err
	}

	return index, nil
}
