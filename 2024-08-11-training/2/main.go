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
		var productsCnt int
		var rate, comissionSum float64
		fmt.Fscan(in, &productsCnt, &rate)

		for range productsCnt {
			var price float64
			fmt.Fscan(in, &price)

			var comission = price * rate / 100.0

			comission = comission - float64(int(comission))

			comissionSum += comission
		}

		fmt.Fprintf(out, "%.2f\n", comissionSum)
	}
}
