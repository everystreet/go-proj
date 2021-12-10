package proj_test

import (
	"testing"

	"github.com/everystreet/go-proj/v8/cproj"
	"github.com/everystreet/go-proj/v8/proj"
	"github.com/stretchr/testify/require"
)

func TestXY(t *testing.T) {
	coord := proj.XY{
		X: 2.52,
		Y: 49.78,
	}

	var cCoord cproj.PJ_COORD
	coord.PutCoordinate(&cCoord)

	var newCoord proj.XY
	newCoord.FromCoordinate(cCoord)
	require.Equal(t, coord, newCoord)
}

func TestXYZ(t *testing.T) {
	coord := proj.XYZ{
		X: 2.52,
		Y: 49.78,
		Z: 10.00,
	}

	var cCoord cproj.PJ_COORD
	coord.PutCoordinate(&cCoord)

	var newCoord proj.XYZ
	newCoord.FromCoordinate(cCoord)
	require.Equal(t, coord, newCoord)
}

func TestXYZT(t *testing.T) {
	coord := proj.XYZT{
		XYZ: proj.XYZ{
			X: 2.52,
			Y: 49.78,
			Z: 10.00,
		},
		T: 120.60,
	}

	var cCoord cproj.PJ_COORD
	coord.PutCoordinate(&cCoord)

	var newCoord proj.XYZT
	newCoord.FromCoordinate(cCoord)
	require.Equal(t, coord, newCoord)
}

func TestLP(t *testing.T) {
	coord := proj.LP{
		Lng: 0.78,
		Lat: 2.57,
	}

	var cCoord cproj.PJ_COORD
	coord.PutCoordinate(&cCoord)

	var newCoord proj.LP
	newCoord.FromCoordinate(cCoord)
	require.Equal(t, coord, newCoord)
}

func TestLPZ(t *testing.T) {
	coord := proj.LPZ{
		LP: proj.LP{
			Lng: 0.78,
			Lat: 2.57,
		},
		Z: 10.00,
	}

	var cCoord cproj.PJ_COORD
	coord.PutCoordinate(&cCoord)

	var newCoord proj.LPZ
	newCoord.FromCoordinate(cCoord)
	require.Equal(t, coord, newCoord)
}

func TestLPZT(t *testing.T) {
	coord := proj.LPZT{
		LPZ: proj.LPZ{
			LP: proj.LP{
				Lng: 0.78,
				Lat: 2.57,
			},
			Z: 10.00,
		},
		T: 120.60,
	}

	var cCoord cproj.PJ_COORD
	coord.PutCoordinate(&cCoord)

	var newCoord proj.LPZT
	newCoord.FromCoordinate(cCoord)
	require.Equal(t, coord, newCoord)
}
