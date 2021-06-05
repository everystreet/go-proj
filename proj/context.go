package proj

import (
	"fmt"
	"github.com/everystreet/go-proj/v6/cproj"
)

type Context struct {
	context *cproj.PJ_CONTEXT

	src        *cproj.PJ
	dst        *cproj.PJ
	pj         *cproj.PJ
	normalized *cproj.PJ
}

func CreateContext(source, target string) (*Context, error) {
	ctx := &Context{}
	ctx.context = cproj.Context_create()

	src, err := CRS(source).instantiate(ctx.context)
	if err != nil {
		defer ctx.Destroy()
		return nil, fmt.Errorf("invalid source '%v': %w", source, err)
	}
	ctx.src = src

	dst, err := CRS(target).instantiate(ctx.context)
	if err != nil {
		defer ctx.Destroy()
		return nil, fmt.Errorf("invalud target '%v': %w", target, err)
	}
	ctx.dst = dst

	ctx.pj = cproj.Create_crs_to_crs_from_pj(ctx.context, src, dst, nil, nil)
	if ctx.pj == nil {
		defer ctx.Destroy()
		return nil, fmt.Errorf("invalid source or target CRS")
	}

	ctx.normalized = cproj.Normalize_for_visualization(ctx.context, ctx.pj)
	if ctx.normalized == nil {
		defer ctx.Destroy()
		return nil, fmt.Errorf("failed to normalize")
	}

	return ctx, nil
}

func (c *Context) Destroy() {
	if c.context != nil {
		cproj.Context_destroy(c.context)
	}
	if c.src != nil {
		cproj.Destroy(c.src)
	}
	if c.dst != nil {
		cproj.Destroy(c.dst)
	}
	if c.pj != nil {
		cproj.Destroy(c.pj)
	}
	if c.normalized != nil {
		cproj.Destroy(c.normalized)
	}
}

func (c *Context) TransformForward(coord Coordinate) {
	transform(c.normalized, forward, coord)
}

func (c *Context) TransformInverse(coord Coordinate) {
	transform(c.normalized, inverse, coord)
}
