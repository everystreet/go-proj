# go-proj

Go bindings for PROJ (formerly PROJ.4) - a library for performing conversions between cartographic projections, maintained by the OSGeo organization. See https://proj.org for more information.

## Usage

`go-proj` uses Cgo to wrap the PROJ library with a Go interface. As such, in order to use this package, you must first install PROJ. The PROJ website contains [details several methods](https://proj.org/install.html) for installing the library, but you may find that these sources are not up-to-date enough to provide the version that this package targets (see `PROJ_VERSION` in [Dockerfile](/Dockerfile)) - in this case you can either compile from source (ensuring that you're using the correct tag), or use the [official Docker images](https://hub.docker.com/r/osgeo/proj/).

You can verify the installation by running the command below and checking that the version number matches the aforementioned `PROJ_VERSION`:

```bash
proj --version
```

Then `go-proj` can be added to your Go module as usual:

```bash
go get -u github.com/everystreet/go-proj
```

`go-proj` consists of two packages - `cproj` contains raw Go bindings to the PROJ library, and `proj` consists of a more idiomatic Go wrapper that doesn't require manual memory deallocation. The latter package, `proj` should be used where possible by adding the following `import` line:

```go
import "github.com/everystreet/go-proj/proj"
```

## Examples

Below is an example of transforming a coordinate to the "Web Mercator" projection:

```go
coord := proj.XYZ{
    X: 2,
    Y: 49,
    Z: 10,
}

err := proj.CRSToCRS(
    proj.CRS("+proj=latlong"),
    proj.CRS("EPSG:3857"),
    proj.TransformForward(&coord))

fmt.Printf("%.2f %.2f", coord.X, coord.Y) // 222638.98 6274861.39
```

`CRSToCRS` is a wrapper around several PROJ functions, including `proj_create_crs_to_crs_from_pj`. The source and target CRS values are passed to this function in the same way as the `cs2cs` program, so the snippet above produces the same result as running the command below:

```bash
$ echo 2 49 | cs2cs +proj=latlong +to EPSG:3857
# 222638.98 6274861.39 0.00
```

For demonstration purposes, this repository also contains a simplified `cs2cs` program that can be run like so:

```bash
$ echo 2 49 | go run cmd/proj/main.go +proj=latlong +to EPSG:3857
# 222638.98 6274861.39 0.00
```

The Go code above is equavalent to the following C snippet (based on the ["Quick start" guide](https://proj.org/development/quickstart.html)):

```c
PJ_CONTEXT *C;
PJ *P;
PJ *P_for_GIS;
PJ_COORD a, b;

C = proj_context_create();

P = proj_create_crs_to_crs(C, "+proj=latlong", "EPSG:3857", NULL);

P_for_GIS = proj_normalize_for_visualization(C, P);
proj_destroy(P);
P = P_for_GIS;

a = proj_coord(2, 49, 10, 0);

b = proj_trans(P, PJ_FWD, a);

proj_destroy(P);
proj_context_destroy(C);
```

## Developing

These bindings are generated by "[c-for-go](https://c.for-go.com)". By running the commands below, the bindings that make up the `cproj` package can be regenerated:

```bash
podman build -t go-proj-builder proj
podman run -v ${PWD}:/build:Z --rm go-proj-builder
```
