package proj

import (
	"github.com/golang/geo/s2"
)

// Area of use for a particular CRS.
type Area struct {
	BottomLeft s2.LatLng
	TopRight   s2.LatLng
}
