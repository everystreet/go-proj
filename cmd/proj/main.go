package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/everystreet/go-proj/v8/proj"
)

func main() {
	var source, target string
	to := false
	for _, arg := range os.Args[1:] {
		switch {
		case arg == "+to":
			to = true
		case to:
			target += " " + arg
		default:
			source += " " + arg
		}
	}

	source = strings.TrimSpace(source)
	target = strings.TrimSpace(target)

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

		if err := proj.CRSToCRS(source, target, func(pj proj.Projection) {
			proj.TransformForward(pj, &coord)
		}); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Printf("%.2f\t%.2f\t%.2f\n", coord.X, coord.Y, coord.Z)
	}
}
