package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var cnt int
	fmt.Fscan(in, &cnt)

	for i := 0; i < cnt; i++ {
		var numbersCnt int
		fmt.Fscan(in, &numbersCnt)

		numbers := make([]int, numbersCnt)

		for j := 0; j < numbersCnt; j++ {
			var number int
			fmt.Fscan(in, &number)

			numbers[j] = number
		}

		fmt.Fprintln(out, getMirrors(numbers))
	}
}

func getMirrors(numbers []int) int {
	result := 0

	for i := 1; i < len(numbers)-2; i++ {
		for j := i + 1; j < len(numbers)-1; j++ {
			if numbers[i]+numbers[j] == numbers[i-1]+numbers[j+1] {
				result++
			}
		}
	}

	return result
}

func getMirrorsV2(numbers []int) int {
	result := 0

	mapDiffsIndexes := make(map[int][]int)

	for j := 1; j < len(numbers)-2; j++ {
		diff := numbers[j] - numbers[j-1]

		mapDiffsIndexes[diff] = append(mapDiffsIndexes[diff], j)
	}

	for j := len(numbers) - 2; j >= 2; j-- {
		diff := numbers[j+1] - numbers[j]

		if diffIndexes, ok := mapDiffsIndexes[diff]; ok {
			for _, indexVal := range diffIndexes {
				if j <= indexVal {
					break
				}

				result++
			}
		}
	}

	return result
}
