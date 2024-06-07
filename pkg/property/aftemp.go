package property

import (
	wtfms "github.com/MeKo-Tech/wtfms/pkg"
)

type AfTemperature struct {
	Temperature
}

func (*AfTemperature) Name() string {
	return "af_temperature"
}

func init() {
	wtfms.RegisterProperty(&AfTemperature{})
}
