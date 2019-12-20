package proj_test

import (
	"testing"

	"github.com/everystreet/go-proj/proj"
	"github.com/stretchr/testify/require"
)

func TestCRSString(t *testing.T) {
	crs := proj.CRS("proj=merc")
	require.Equal(t, "proj=merc +type=crs", crs.String())
}
