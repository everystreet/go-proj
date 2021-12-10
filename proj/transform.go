package proj

import (
	"fmt"

	"github.com/everystreet/go-proj/v8/cproj"
)

// Projection context.
type Projection *cproj.PJ

// CRSToCRS facilitates transformations between two coordinate reference systems.
func CRSToCRS(source, target string, transform func(Projection)) error {
	ctx := cproj.Context_create()
	defer cproj.Context_destroy(ctx)

	src, err := CRS(source).instantiate(ctx)
	if err != nil {
		return fmt.Errorf("invalid source '%v': %w", source, err)
	}
	defer cproj.Destroy(src)

	dst, err := CRS(target).instantiate(ctx)
	if err != nil {
		return fmt.Errorf("invalud target '%v': %w", target, err)
	}
	defer cproj.Destroy(dst)

	pj := cproj.Create_crs_to_crs_from_pj(ctx, src, dst, nil, nil)
	if pj == nil {
		return fmt.Errorf("invalid source or target CRS")
	}
	defer cproj.Destroy(pj)

	normalized := cproj.Normalize_for_visualization(ctx, pj)
	if normalized == nil {
		return fmt.Errorf("failed to normalize")
	}
	defer cproj.Destroy(normalized)

	transform(normalized)
	return nil
}

// Coordinate wraps functions that allow communication with the cproj package.
type Coordinate interface {
	PutCoordinate(*cproj.PJ_COORD)
	FromCoordinate(cproj.PJ_COORD)
}

// TransformForward performs a forward transformation of the supplied coordinate.
func TransformForward(pj Projection, coord Coordinate) {
	transform(pj, forward, coord)
}

// TransformInverse performs an inverse transformation of the supplied coordinate.
func TransformInverse(pj Projection, coord Coordinate) {
	transform(pj, inverse, coord)
}

const (
	forward cproj.PJ_DIRECTION = 1
	inverse cproj.PJ_DIRECTION = -1
)

func transform(pj *cproj.PJ, direction cproj.PJ_DIRECTION, coord Coordinate) {
	var in cproj.PJ_COORD
	coord.PutCoordinate(&in)

	out := cproj.Trans(pj, direction, in)
	coord.FromCoordinate(out)
}
