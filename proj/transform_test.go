package proj_test

import (
	"fmt"
	"testing"

	"github.com/everystreet/go-proj/v6/proj"
)

func ExampleCRSToCRS() {
	coord := proj.XYZ{
		X: 2,
		Y: 49,
		Z: 10,
	}

	if err := proj.CRSToCRS("+proj=latlong", "EPSG:3857", func(pj proj.Projection) {
		proj.TransformForward(pj, &coord)
		// transform more coordinates
	}); err != nil {
		panic(err)
	}

	fmt.Printf("%.2f %.2f %.2f", coord.X, coord.Y, coord.Z)
	// Output: 222638.98 6274861.39 10.00
}

func BenchmarkCRSToCRS(b *testing.B) {
	for i := 0; i < b.N; i++ {
		coord := proj.XYZ{
			X: 2,
			Y: 49,
			Z: 10,
		}

		if err := proj.CRSToCRS("+proj=latlong", "EPSG:3857", func(pj proj.Projection) {
			proj.TransformForward(pj, &coord)
			// transform more coordinates
		}); err != nil {
			panic(err)
		}
	}
}

func BenchmarkContext(b *testing.B) {
	var context, _ = proj.CreateContext("+proj=latlong", "EPSG:3857")
	defer context.Destroy()

	for i := 0; i < b.N; i++ {
		coord := proj.XYZ{
			X: 2,
			Y: 49,
			Z: 10,
		}

		context.TransformInverse(&coord)
	}
}
