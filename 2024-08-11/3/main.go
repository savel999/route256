package main

import (
	"bufio"
	"fmt"
	"os"
)

type ProductData struct {
	name       string
	start, end int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var cnt int
	fmt.Fscan(in, &cnt)

	for i := 0; i < cnt; i++ {
		var requestsCnt int
		fmt.Fscan(in, &requestsCnt)

		productsHistoryMap := make(map[int][]ProductData)
		productNameIdMap := make(map[string]int)

		requestTime := 0

		for j := 0; j < requestsCnt; j++ {
			requestTime++

			var action string
			fmt.Fscan(in, &action)

			if action == "CHANGE" {
				var (
					name string
					id   int
				)

				fmt.Fscan(in, &name, &id)

				if _, ok := productsHistoryMap[id]; !ok {
					productsHistoryMap[id] = make([]ProductData, 0)
				}

				if oldId, ok := productNameIdMap[name]; ok {
					historyList := productsHistoryMap[oldId]
					productsHistoryMap[oldId][len(historyList)-1].end += requestTime
				}

				// новый id-шник
				if len(productsHistoryMap[id]) > 0 {
					productsHistoryMap[id][len(productsHistoryMap[id])-1].end = requestTime
				}

				productsHistoryMap[id] = append(productsHistoryMap[id], ProductData{name: name, start: requestTime})
				productNameIdMap[name] = id

			} else {
				var (
					searchId, time int
				)

				fmt.Fscan(in, &searchId, &time)

				productHistory := productsHistoryMap[searchId]
				foundProductName, isFinded := findProductIdFromHistoryByTime(productHistory, time)

				if isFinded {
					fmt.Fprintln(out, foundProductName)
				} else {
					fmt.Fprintln(out, "404")
				}
			}
		}
	}
}

func findProductIdFromHistoryByTime(history []ProductData, time int) (string, bool) {
	if len(history) == 0 {
		return "", false
	}

	for _, historyItem := range history {
		if time >= historyItem.start && (time < historyItem.end || historyItem.end == 0) {
			return historyItem.name, true
		}
	}

	return "", false
}
