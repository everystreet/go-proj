package proj

import (
	"fmt"
	"strings"

	"github.com/everystreet/go-proj/cproj"
)

// CRS is a coordinate reference system definition.
type CRS string

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
