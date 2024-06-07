package property

import (
	"encoding/json"
	"fmt"

	wtfms "github.com/MeKo-Tech/wtfms/pkg"
)

type TubeDiameter struct {
	Inner Diameter `json:"inner"`
	Outer Diameter `json:"outer"`
}

func (*TubeDiameter) Name() string {
	return "tube_diameter"
}

func (t *TubeDiameter) Validate(data json.RawMessage) error {
	// TODO: use reflection to instantiate the data
	var diameter TubeDiameter
	err := json.Unmarshal(data, &diameter)
	if err != nil {
		return err
	}

	if diameter.Inner.Nominal.GreaterThanOrEqual(diameter.Outer.Nominal.Decimal) {
		return fmt.Errorf("inner diameter is larger than outer diameter")
	}

	return nil
}

func init() {
	wtfms.RegisterProperty(&TubeDiameter{})
}
