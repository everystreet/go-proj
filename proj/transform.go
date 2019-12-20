package proj

import (
	"fmt"

	"github.com/everystreet/go-proj/cproj"
)

// Transformation functions are passed to CRSToCRS.
type Transformation func(*cproj.PJ)

// CRSToCRS facilitates transformations between two coordinate reference systems.
func CRSToCRS(source, target CRS, ts ...Transformation) error {
	ctx := cproj.Context_create()
	defer cproj.Context_destroy(ctx)

	src, err := source.instantiate(ctx)
	if err != nil {
		return fmt.Errorf("invalid source '%v': %w", source, err)
	}
	defer cproj.Destroy(src)

	dst, err := target.instantiate(ctx)
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

	for _, t := range ts {
		t(normalized)
	}
	return nil
}

// Coordinate wraps functions that allow communication with the cproj package.
type Coordinate interface {
	PutCoordinate(*cproj.PJ_COORD)
	FromCoordinate(cproj.PJ_COORD)
}

// TransformForward performs a forward transformation of the supplied coordinate.
func TransformForward(coord Coordinate) Transformation {
	return transform(forward, coord)
}

// TransformInverse performs an inverse transformation of the supplied coordinate.
func TransformInverse(coord Coordinate) Transformation {
	return transform(inverse, coord)
}

func transform(direction cproj.PJ_DIRECTION, coord Coordinate) Transformation {
	return func(pj *cproj.PJ) {
		var in cproj.PJ_COORD
		coord.PutCoordinate(&in)

		out := cproj.Trans(pj, direction, in)
		coord.FromCoordinate(out)
	}
}

const (
	forward cproj.PJ_DIRECTION = 1
	inverse cproj.PJ_DIRECTION = -1
)
