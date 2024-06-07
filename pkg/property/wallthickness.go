package property

import (
	wtfms "github.com/MeKo-Tech/wtfms/pkg"
)

type WallThickness struct {
	Length
}

func (*WallThickness) Name() string {
	return "wall_thickness"
}

func init() {
	wtfms.RegisterProperty(&WallThickness{})
}
