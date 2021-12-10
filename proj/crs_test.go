package proj_test

import (
	"testing"

	"github.com/everystreet/go-proj/v8/proj"
	"github.com/golang/geo/s2"
	"github.com/stretchr/testify/require"
)

func TestCRSString(t *testing.T) {
	crs := proj.CRS("proj=merc")
	require.Equal(t, "proj=merc +type=crs", crs.String())
}

func TestCRSArea(t *testing.T) {
	crs := proj.CRS("EPSG:3857")

	area, ok, err := crs.Area()
	require.NoError(t, err)
	require.True(t, ok)
	require.Equal(t, proj.Area{
		BottomLeft: s2.LatLngFromDegrees(-85.06, -180),
		TopRight:   s2.LatLngFromDegrees(85.06, 180),
	}, *area)
}
