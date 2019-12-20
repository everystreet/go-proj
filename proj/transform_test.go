package proj_test

import (
	"fmt"

	"github.com/everystreet/go-proj/proj"
)

func ExampleCRSToCRS() {
	coord := proj.XYZ{
		X: 2,
		Y: 49,
		Z: 10,
	}

	if err := proj.CRSToCRS(proj.CRS("+proj=latlong"), proj.CRS("EPSG:3857"), proj.TransformForward(&coord)); err != nil {
		panic(err)
	}

	fmt.Printf("%.2f %.2f %.2f", coord.X, coord.Y, coord.Z)
	// Output: 222638.98 6274861.39 10.00
}
