package main

import (
	"fmt"

	"github.com/everystreet/go-proj/cproj"
)

func main() {
	fmt.Println("test ll safestrings")
	//spew.Dump(bindings.Info())

	ctx := cproj.Context_create()
	fmt.Println(cproj.Context_get_database_path(ctx))

	fmt.Println(cproj.Create(ctx, "proj=latlong"))

	//fmt.Println(bindings.Context_set_database_path(nil, "/usr/local/share/proj/proj.db", nil, nil))
	//bindings.Create_crs_to_crs(nil, "EPSG:25832", "EPSG:25833", nil)
	fmt.Println("done")
}
