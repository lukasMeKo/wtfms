package types

import (
	"fmt"
	"strings"

	"github.com/invopop/jsonschema"
	"github.com/shopspring/decimal"
)

type Decimal struct{ decimal.Decimal }

func (d *Decimal) Unwrap() *decimal.Decimal {
	return &d.Decimal
}

func (d *Decimal) UnmarshalJSON(data []byte) error {
	dec, err := decimal.NewFromString(strings.Trim(string(data), `"`))
	if err != nil {
		return fmt.Errorf("failed to parse '%s' as decimal number: %w", data, err)
	}
	*d = Decimal{dec}

	return nil
}

func (d *Decimal) MarshalJSON() ([]byte, error) {
	return []byte(d.Unwrap().String()), nil
}

func (Decimal) JSONSchema() *jsonschema.Schema {
	return &jsonschema.Schema{
		Type:        "number",
		Title:       "decimal number",
		Description: "",
	}
}
