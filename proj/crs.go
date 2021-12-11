package proj

import (
	"fmt"
	"strings"

	"github.com/everystreet/go-proj/v8/cproj"
	"github.com/golang/geo/s2"
)

// CRS is a coordinate reference system definition.
type CRS string

// Area returns the area of use for the CRS, or false if no such area is defined.
func (c CRS) Area() (*Area, bool, error) {
	ctx := cproj.Context_create()
	defer cproj.Context_destroy(ctx)

	pj, err := c.instantiate(ctx)
	if err != nil {
		return nil, false, fmt.Errorf("invalid source '%v': %w", c, err)
	}
	defer cproj.Destroy(pj)

	var north, east, south, west float64
	if cproj.Get_area_of_use(ctx, pj, &west, &south, &east, &north, nil) != 1 {
		return nil, false, nil
	}

	if north == -1000 || east == -1000 || south == -1000 || west == -1000 {
		return nil, false, nil
	}

	return &Area{
		BottomLeft: s2.LatLngFromDegrees(south, west),
		TopRight:   s2.LatLngFromDegrees(north, east),
	}, true, nil
}

// Area of use for a particular CRS.
type Area struct {
	BottomLeft s2.LatLng
	TopRight   s2.LatLng
}

// String returns the lower-cased CRS definition.
// If the definition is a proj-string, the "+type=crs" option if appended if not present.
func (c CRS) String() string {
	str := strings.ToLower(string(c))
	if strings.HasPrefix(str, "proj=") || strings.HasPrefix(str, "+proj=") ||
		strings.HasPrefix(str, "+init=") || strings.HasPrefix(str, "+title=") &&
		!strings.Contains(str, "type=crs") {
		str += " +type=crs"
	}
	return str
}

func (c CRS) instantiate(ctx *cproj.PJ_CONTEXT) (*cproj.PJ, error) {
	pj := cproj.Create(ctx, c.String())
	if pj == nil {
		return nil, fmt.Errorf("failed to create CRS")
	}
	return pj, nil
}
