package proj

import (
	"github.com/everystreet/go-proj/v8/cproj"
	"github.com/golang/geo/r2"
	"github.com/golang/geo/r3"
	"github.com/golang/geo/s1"
	"github.com/golang/geo/s2"
)

// XY is a 2D cartesian coordinate.
type XY r2.Point

// PutCoordinate updates coord with the values in c.
func (c XY) PutCoordinate(coord *cproj.PJ_COORD) {
	for data, i := float64ToDoubleBytes(c.X), 0; i < 8; i++ {
		coord[i+0] = data[i]
	}

	for data, i := float64ToDoubleBytes(c.Y), 0; i < 8; i++ {
		coord[i+8] = data[i]
	}
}

// FromCoordinate updates c with the values in coord.
func (c *XY) FromCoordinate(coord cproj.PJ_COORD) {
	c.X = doubleBytesToFloat64(coord[0:8])
	c.Y = doubleBytesToFloat64(coord[8:16])
}

// XYZ is a 3D cartesian coordinate.
type XYZ r3.Vector

// PutCoordinate updates coord with the values in c.
func (c XYZ) PutCoordinate(coord *cproj.PJ_COORD) {
	for data, i := float64ToDoubleBytes(c.X), 0; i < 8; i++ {
		coord[i+0] = data[i]
	}

	for data, i := float64ToDoubleBytes(c.Y), 0; i < 8; i++ {
		coord[i+8] = data[i]
	}

	for data, i := float64ToDoubleBytes(c.Z), 0; i < 8; i++ {
		coord[i+16] = data[i]
	}
}

// FromCoordinate updates c with the values in coord.
func (c *XYZ) FromCoordinate(coord cproj.PJ_COORD) {
	c.X = doubleBytesToFloat64(coord[0:8])
	c.Y = doubleBytesToFloat64(coord[8:16])
	c.Z = doubleBytesToFloat64(coord[16:24])
}

// XYZT is a 3D cartesian coordinate with a time component.
type XYZT struct {
	XYZ
	T float64
}

// PutCoordinate updates coord with the values in c.
func (c XYZT) PutCoordinate(coord *cproj.PJ_COORD) {
	c.XYZ.PutCoordinate(coord)
	for data, i := float64ToDoubleBytes(c.T), 0; i < 8; i++ {
		coord[i+24] = data[i]
	}
}

// FromCoordinate updates c with the values in coord.
func (c *XYZT) FromCoordinate(coord cproj.PJ_COORD) {
	c.XYZ.FromCoordinate(coord)
	c.T = doubleBytesToFloat64(coord[24:32])
}

// LP is a geodetic coordinate expressed in radians.
type LP s2.LatLng

// PutCoordinate updates coord with the values in c.
func (c LP) PutCoordinate(coord *cproj.PJ_COORD) {
	for data, i := float64ToDoubleBytes(c.Lng.Radians()), 0; i < 8; i++ {
		coord[i+0] = data[i]
	}

	for data, i := float64ToDoubleBytes(c.Lat.Radians()), 0; i < 8; i++ {
		coord[i+8] = data[i]
	}
}

// FromCoordinate updates c with the values in coord.
func (c *LP) FromCoordinate(coord cproj.PJ_COORD) {
	c.Lng = s1.Angle(doubleBytesToFloat64(coord[0:8]))
	c.Lat = s1.Angle(doubleBytesToFloat64(coord[8:16]))
}

// LPZ is a geodetic coordinate expressed in radians, with a vertical component.
type LPZ struct {
	LP
	Z float64
}

// PutCoordinate updates coord with the values in c.
func (c LPZ) PutCoordinate(coord *cproj.PJ_COORD) {
	c.LP.PutCoordinate(coord)
	for data, i := float64ToDoubleBytes(c.Z), 0; i < 8; i++ {
		coord[i+16] = data[i]
	}
}

// FromCoordinate updates c with the values in coord.
func (c *LPZ) FromCoordinate(coord cproj.PJ_COORD) {
	c.LP.FromCoordinate(coord)
	c.Z = doubleBytesToFloat64(coord[16:24])
}

// LPZT is a geodetic coordinate expressed in radians, with vertical and time components.
type LPZT struct {
	LPZ
	T float64
}

// PutCoordinate updates coord with the values in c.
func (c LPZT) PutCoordinate(coord *cproj.PJ_COORD) {
	c.LPZ.PutCoordinate(coord)
	for data, i := float64ToDoubleBytes(c.T), 0; i < 8; i++ {
		coord[i+24] = data[i]
	}
}

// FromCoordinate updates c with the values in coord.
func (c *LPZT) FromCoordinate(coord cproj.PJ_COORD) {
	c.LPZ.FromCoordinate(coord)
	c.T = doubleBytesToFloat64(coord[24:32])
}
