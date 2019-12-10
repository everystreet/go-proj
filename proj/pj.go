package proj

import (
	"github.com/everystreet/go-proj/cproj"
)

type Transformation func(*cproj.PJ)

func CRSToCRS(source, target string, ts ...Transformation) {
	ctx := cproj.Context_create()
	defer cproj.Context_destroy(ctx)

	pj := cproj.Create_crs_to_crs(ctx, source, target, nil)
	defer cproj.Destroy(pj)

	normalized := cproj.Normalize_for_visualization(ctx, pj)
	defer cproj.Destroy(normalized)

	for _, t := range ts {
		t(normalized)
	}
}

type SourceCoordinate interface {
	PutCoordinate(*cproj.PJ_COORD)
}

type DestinationCoordinate interface {
	FromCoordinate(cproj.PJ_COORD)
}

func TransformForward(coord SourceCoordinate, dest DestinationCoordinate) Transformation {
	return transform(coord, forward, dest)
}

func TransformInverse(coord SourceCoordinate, dest DestinationCoordinate) Transformation {
	return transform(coord, inverse, dest)
}

func transform(coord SourceCoordinate, direction cproj.PJ_DIRECTION, dest DestinationCoordinate) Transformation {
	return func(pj *cproj.PJ) {
		var in cproj.PJ_COORD
		coord.PutCoordinate(&in)

		out := cproj.Trans(pj, direction, in)
		dest.FromCoordinate(out)
	}
}

const (
	forward cproj.PJ_DIRECTION = 1
	inverse cproj.PJ_DIRECTION = -1
)
