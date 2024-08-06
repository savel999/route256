package main

import (
	"bufio"
	"fmt"
	"os"
)

type CityResources struct {
	Width, Height int
	Resources     [][]Point
}

type Point struct {
	X, Y int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var cnt int
	fmt.Fscan(in, &cnt)

	for i := 0; i < cnt; i++ {
		var width, height, resourcesTypes int
		fmt.Fscan(in, &width, &height, &resourcesTypes)

		cityResources := CityResources{
			Width:     width,
			Height:    height,
			Resources: make([][]Point, resourcesTypes),
		}

		for j := 0; j < resourcesTypes; j++ {
			var resourceCnt int
			fmt.Fscan(in, &resourceCnt)

			cityResources.Resources[j] = make([]Point, resourceCnt)

			for k := 0; k < resourceCnt; k++ {
				var x, y int
				fmt.Fscan(in, &x, &y)

				cityResources.Resources[j][k] = Point{X: x, Y: y}
			}
		}

		fmt.Fprintln(out, calculateMinSquare(cityResources))
	}
}

func calculateMinSquare(r CityResources) int {
	// определяем стартовые точки
	startResourceIndex, resourceCnt := 0, len(r.Resources[0])

	for i := 1; i < len(r.Resources); i++ {
		if len(r.Resources[i]) < resourceCnt {
			resourceCnt = len(r.Resources[i])
			startResourceIndex = i
		}
	}

	startsPoints, minSquare := r.Resources[startResourceIndex], 0

	for _, point := range startsPoints {
		// посчитать площадь
	}

	return minSquare
}
