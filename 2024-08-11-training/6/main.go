package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

type TransportInfo struct {
	Cars, Capacity int
	Boxes          []Box
}

type Box struct {
	Weight, Count int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var cnt int
	fmt.Fscan(in, &cnt)

	for i := 0; i < cnt; i++ {
		var carsCnt, capacity int
		fmt.Fscan(in, &carsCnt, &capacity)

		var boxesCnt int
		fmt.Fscan(in, &boxesCnt)

		mapBoxes := make(map[int]int)

		for j := 0; j < boxesCnt; j++ {
			var boxWeight int
			fmt.Fscan(in, &boxWeight)

			mapBoxes[int(math.Pow(2.0, float64(boxWeight)))]++
		}

		boxes := make([]Box, 0, len(mapBoxes))

		for boxWeight, count := range mapBoxes {
			boxes = append(boxes, Box{Weight: boxWeight, Count: count})
		}

		sort.Slice(boxes, func(i, j int) bool { return boxes[i].Weight > boxes[j].Weight })

		fmt.Fprintln(
			out,
			getTransportsCount(TransportInfo{Cars: carsCnt, Capacity: capacity, Boxes: boxes}),
		)
	}
}

func getTransportsCount(info TransportInfo) int {
	carsCnt := 0

	for len(info.Boxes) > 0 {
		carCapacity := info.Capacity
		newBoxes := make([]Box, 0, len(info.Boxes))

		i := 0
		for i < len(info.Boxes) {
			if carCapacity >= info.Boxes[i].Weight {
				carCapacity -= info.Boxes[i].Weight

				info.Boxes[i].Count--
				if info.Boxes[i].Count == 0 {
					i++
				}

				continue
			}

			newBoxes = append(newBoxes, info.Boxes[i])
			i++
		}

		info.Boxes = newBoxes
		carsCnt++
	}

	return int(math.Ceil((float64(carsCnt) / float64(info.Cars))))
}
