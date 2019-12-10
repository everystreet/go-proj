package proj

import "C"
import (
	"encoding/binary"
	"math"

	"unsafe"

	"github.com/everystreet/go-proj/cproj"
	"github.com/golang/geo/r2"
	"github.com/golang/geo/r3"
	"github.com/golang/geo/s2"
)

// XY is a 2D cartesian coordinate.
type XY r2.Point

func (c XY) PutCoordinate(coord cproj.PJ_COORD) {
	binary.LittleEndian.PutUint64(coord[0:], math.Float64bits(c.X))
	binary.LittleEndian.PutUint64(coord[8:], math.Float64bits(c.Y))
}

// XYZ is a 3D cartesian coordinate.
type XYZ r3.Vector

func (c XYZ) PutCoordinate(coord *cproj.PJ_COORD) {
	d := C.double(c.X)
	for data, i := (*[8]byte)(unsafe.Pointer(&d)), 0; i < 8; i++ {
		coord[i+0] = data[i]
	}

	d = C.double(c.Y)
	for data, i := (*[8]byte)(unsafe.Pointer(&d)), 0; i < 8; i++ {
		coord[i+8] = data[i]
	}

	d = C.double(c.Z)
	for data, i := (*[8]byte)(unsafe.Pointer(&d)), 0; i < 8; i++ {
		coord[i+16] = data[i]
	}
}

func (c *XYZ) FromCoordinate(coord cproj.PJ_COORD) {
	c.X = math.Float64frombits(binary.LittleEndian.Uint64(coord[0:]))
	c.Y = math.Float64frombits(binary.LittleEndian.Uint64(coord[8:]))
	c.Z = math.Float64frombits(binary.LittleEndian.Uint64(coord[16:]))
}

// XYZT is a 3D cartesian coordinate with a time component.
type XYZT struct {
	XYZ
	T float64
}

/*func (c XYZT) PutCoordinate(coord cproj.PJ_COORD) {
	c.XYZ.PutCoordinate(coord)
	//binary.LittleEndian.PutUint64(coord[24:], math.Float64bits(c.T))
}*/

// LP is a geodetic coordinate.
type LP s2.LatLng

// LPZ is a geodetic coordinate with a vertical component.
type LPZ struct {
	LP
	Z float64
}

// LPZT is a geodetic coordinate with vertical and time components.
type LPZT struct {
	LPZ
	T float64
}
