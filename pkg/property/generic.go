package property

import (
	wtfms "github.com/MeKo-Tech/wtfms/pkg"
	"github.com/MeKo-Tech/wtfms/pkg/property/types"
)

type (
	Length struct {
		Nominal   types.Decimal `json:"nominal"`
		Deviation types.Decimal `json:"deviation"`
	}
	Diameter struct {
		Nominal   types.Decimal `json:"nominal"`
		Deviation types.Decimal `json:"deviation"`
	}
	Temperature struct {
		Nominal   types.Decimal `json:"nominal"`
		Deviation types.Decimal `json:"deviation"`
	}
	String struct {
		Value string `json:"value"`
	}
	Decimal struct {
		Value types.Decimal `json:"value"`
	}
)

func init() {
	wtfms.RegisterProperty(&Length{})
	wtfms.RegisterProperty(&Diameter{})
	wtfms.RegisterProperty(&Temperature{})
	wtfms.RegisterProperty(&String{})
	wtfms.RegisterProperty(&Decimal{})
}

// Name implements wtfms.Property.
func (d *Decimal) Name() string {
	return "decimal"
}

// Name implements wtfms.Property.
func (s *String) Name() string {
	return "string"
}

// Name implements wtfms.Property.
func (t *Temperature) Name() string {
	return "temperature"
}

// Name implements wtfms.Property.
func (l *Length) Name() string {
	return "length"
}

// Name implements wtfms.Property.
func (d *Diameter) Name() string {
	return "diameter"
}
