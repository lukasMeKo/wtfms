package wtfms

import (
	"encoding/json"
	"fmt"

	schemagen "github.com/invopop/jsonschema"
	"github.com/meilisearch/meilisearch-go"
	"github.com/qri-io/jsonschema"
)

type Validator interface {
	Validate(data json.RawMessage) error
}

type PropertyType struct {
	NameField   string             `json:"name"`
	SchemaField *jsonschema.Schema `json:"schema"`
}

var _ Property = &PropertyType{}

func (prop *PropertyType) Name() string {
	return prop.NameField
}

// TODO: add hooks for index creation, i.e. filterability etc.
type Property interface {
	Name() string
}

func GetProperty(client *meilisearch.Client, name string) (Property, error) {
	if prop, found := properties[name]; found {
		return prop, nil
	}

	var property PropertyType

	index := client.Index("properties")
	if err := index.GetDocument(name, nil, &property); err != nil {
		return nil, err
	}

	return &property, nil
}

func PropertiesExists(client *meilisearch.Client, names ...string) error {
	index := client.Index("properties")
	for _, name := range names {
		if err := index.GetDocument(name, nil, nil); err != nil {
			return err
		}
	}

	return nil
}

func RegisterProperties(client *meilisearch.Client, properties ...PropertyType) error {
	for _, property := range properties {
		// Could be async!
		if err := ValidateProperty(client, property); err != nil {
			return err
		}
	}

	index := client.Index("properties")
	if _, err := index.AddDocuments(&properties); err != nil {
		return err
	}

	return nil
}

func ValidateProperty(client *meilisearch.Client, property PropertyType) error {
	// schema must be valid json schema
	if property.SchemaField == nil {
		return fmt.Errorf("invalid schema")
	}

	if property.SchemaField.TopLevelType() == "unknown" {
		return fmt.Errorf("invalid schema")
	}

	return nil
}

func schema(v any) *jsonschema.Schema {
	switch value := v.(type) {
	case *PropertyType:
		return value.SchemaField
	default:
		schema := jsonschema.Schema{}
		data, _ := schemagen.Reflect(v).MarshalJSON()
		_ = json.Unmarshal(data, &schema)
		return &schema
	}
}
