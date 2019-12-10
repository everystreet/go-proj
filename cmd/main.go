package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/everystreet/go-proj/proj"
)

func main() {
	for {
		r := bufio.NewReader(os.Stdin)
		str, _ := r.ReadString('\n')
		if str == "" {
			return
		}

		fields := strings.Fields(str)

		getFloat := func(i int) float64 {
			if len(fields) <= i {
				return 0
			}

			f, err := strconv.ParseFloat(fields[i], 10)
			if err != nil {
				os.Exit(1)
			}
			return f
		}

		coord := proj.XYZ{
			X: getFloat(0),
			Y: getFloat(1),
			Z: getFloat(2),
		}

		proj.CRSToCRS("EPSG:4326", "EPSG:32632",
			proj.TransformForward(coord, &coord))

		fmt.Printf("%.2f\t%.2f\t%.2f\n", coord.X, coord.Y, coord.Z)
	}
}
