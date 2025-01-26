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
		var wordsCnt int
		fmt.Fscan(in, &wordsCnt)

		evenParts := make(map[string][]int)
		oddParts := make(map[string][]int)

		pairsMap := make(map[string]bool)

		for j := 0; j < wordsCnt; j++ {
			var wordBytes []byte
			fmt.Fscan(in, &wordBytes)

			word := WordParts{}

			for k, b := range wordBytes {
				if k%2 == 0 {
					word.EvenBytes = append(word.EvenBytes, b)
				} else {
					word.OddBytes = append(word.OddBytes, b)
				}
			}

			evenPart := string(word.EvenBytes)

			if items, ok := evenParts[evenPart]; ok {
				for _, item := range items {
					pairsMap[fmt.Sprintf("%d_%d", item, j)] = true
				}
			}

			evenParts[evenPart] = append(evenParts[evenPart], j)

			if len(word.OddBytes) > 0 {
				oddPart := string(word.OddBytes)

				if items, ok := oddParts[oddPart]; ok {
					for _, item := range items {
						pairsMap[fmt.Sprintf("%d_%d", item, j)] = true
					}
				}

				oddParts[oddPart] = append(oddParts[oddPart], j)
			}
		}

		fmt.Fprintf(out, "%d\n", len(pairsMap))
	}
}

type WordParts struct {
	EvenBytes []byte //четные
	OddBytes  []byte //нечетные
}
