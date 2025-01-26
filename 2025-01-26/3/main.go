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

		evenParts := make(map[string]int)
		oddParts := make(map[string]int)

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

			evenParts[string(word.EvenBytes)]++

			if len(word.OddBytes) > 0 {
				oddParts[string(word.OddBytes)]++
			}
		}

		maxMatches, result := 0, 0

		for _, v := range evenParts {
			maxMatches = max(maxMatches, v)
		}

		for _, v := range oddParts {
			maxMatches = max(maxMatches, v)
		}

		for k := maxMatches - 1; k >= 1; k-- {
			result += k
		}

		fmt.Fprintf(out, "%d\n", result)
	}
}

type WordParts struct {
	EvenBytes []byte //четные
	OddBytes  []byte //нечетные
}
