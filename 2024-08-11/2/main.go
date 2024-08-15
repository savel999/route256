package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var cnt int
	fmt.Fscan(in, &cnt)

	for i := 0; i < cnt; i++ {
		var numberInBytes []byte
		fmt.Fscan(in, &numberInBytes)

		removableIndex := 0

		if len(numberInBytes) == 1 {
			fmt.Fprintln(out, 0)
			break
		} else if len(numberInBytes) == 2 {
			var result2 int

			if numberInBytes[0] > numberInBytes[1] {
				result2, _ = strconv.Atoi(string(numberInBytes[0]))
			} else {
				result2, _ = strconv.Atoi(string(numberInBytes[1]))
			}

			fmt.Fprintln(out, result2)

			break
		}

		for j := 0; j < len(numberInBytes)-1; j++ {
			if numberInBytes[j] <= numberInBytes[j+1] {
				removableIndex = j
				break
			}
		}

		result, _ := strconv.Atoi(string(removeElement(numberInBytes, removableIndex)))

		fmt.Fprintln(out, result)
	}
}

func removeElement(slice []byte, index int) []byte {
	copy(slice[index:], slice[index+1:])
	return slice[:len(slice)-1]
}
