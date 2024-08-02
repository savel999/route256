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
		var requestsCnt int
		fmt.Fscan(in, &requestsCnt)

		requestsList := make([]int, requestsCnt)

		for j := 0; j < requestsCnt; j++ {
			var clientId int
			fmt.Fscan(in, &clientId)

			requestsList[j] = clientId
		}

		left, maxLen, curLen := 0, 0, 0
		clientsNum := 2
		clientsMap := make(map[int]bool, clientsNum)

		for j := left; j < len(requestsList); j++ {
			if _, ok := clientsMap[requestsList[j]]; !ok {
				clientsNum--
			}

			if clientsNum < 0 {
				clientsMap = make(map[int]bool, clientsNum)
				maxLen = max(maxLen, curLen)
				clientsNum = 2
				curLen = 0

				j = left
				left++

				continue
			}

			curLen++
			clientsMap[requestsList[j]] = true
		}

		maxLen = max(maxLen, curLen)

		fmt.Fprintln(out, maxLen)
	}
}
