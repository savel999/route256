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
		var rowsCnt, columnsCnt int
		fmt.Fscan(in, &rowsCnt, &columnsCnt)

		if rowsCnt == 1 || columnsCnt == 1 {
			fmt.Fprintln(out, "1")
			if rowsCnt == 1 {
				fmt.Fprintln(out, "1 1 R")
			} else {
				fmt.Fprintln(out, "1 1 D")
			}

		} else if rowsCnt > columnsCnt {
			fmt.Fprintln(out, "2")
			fmt.Fprintln(out, "1 1 D")
			fmt.Fprintf(out, "%d %d U\n", rowsCnt, columnsCnt)
		} else {
			fmt.Fprintln(out, "2")
			fmt.Fprintln(out, "1 1 R")
			fmt.Fprintf(out, "%d %d L\n", rowsCnt, columnsCnt)
		}

	}
}
