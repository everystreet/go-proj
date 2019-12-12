package proj

import (
	"fmt"

	"github.com/everystreet/go-proj/cproj"
)

type Transformation func(*cproj.PJ)

func CRSToCRS(source, target string, ts ...Transformation) error {
	ctx := cproj.Context_create()
	defer cproj.Context_destroy(ctx)

	pj := cproj.Create_crs_to_crs(ctx, source, target, nil)
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

type Coordinate interface {
	PutCoordinate(*cproj.PJ_COORD)
	FromCoordinate(cproj.PJ_COORD)
}

func TransformForward(coord Coordinate) Transformation {
	return transform(forward, coord)
}

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
