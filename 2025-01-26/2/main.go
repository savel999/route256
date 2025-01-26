package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func main() {
	in := bufio.NewReaderSize(os.Stdin, 4*1024*1024)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var cnt int
	fmt.Fscan(in, &cnt)

	for i := 0; i < cnt; i++ {
		var productsCnt int
		fmt.Fscan(in, &productsCnt)

		mapPriceProducts := make(map[int]map[string]bool)

		for j := 0; j < productsCnt; j++ {
			var name string
			fmt.Fscan(in, &name)
			var price int
			fmt.Fscan(in, &price)

			if _, ok := mapPriceProducts[price]; !ok {
				mapPriceProducts[price] = make(map[string]bool)
			}

			mapPriceProducts[price][name] = false
		}

		var resultListLine []byte
		fmt.Fscan(in, &resultListLine)

		if validate(resultListLine, mapPriceProducts) {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}
}

func validate(resultLine []byte, mapPriceProducts map[int]map[string]bool) bool {
	splitted := bytes.Split(resultLine, []byte{','})

	resultPriceProducts := make(map[int]int)

	for _, part := range splitted {
		productParts := bytes.Split(part, []byte{':'})

		if len(productParts) != 2 {
			return false
		}

		name := string(productParts[0])
		price := productParts[1]

		if len(price) == 0 || price[0] == '0' {
			return false // ведущий 0
		}

		parsedPriceInt, err := strconv.Atoi(string(price))
		if err != nil || parsedPriceInt == 0 {
			return false
		}

		resultPriceProducts[parsedPriceInt]++

		if resultPriceProducts[parsedPriceInt] > 1 {
			return false
		}

		if _, ok := mapPriceProducts[parsedPriceInt][name]; !ok {
			return false
		} else {
			mapPriceProducts[parsedPriceInt][name] = true
		}
	}

	for _, productsList := range mapPriceProducts {
		isPicked := false

		for _, picked := range productsList {
			if picked {
				isPicked = true

				break
			}
		}

		if !isPicked {
			return false
		}
	}

	return true
}
