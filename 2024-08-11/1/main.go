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

		for j := 1; j < len(numberInBytes)-1; j++ {
			//если число посередине меньше левого и правого
			if numberInBytes[j-1] > numberInBytes[j] && numberInBytes[j] < numberInBytes[j+1] {
				removableIndex = j
				break
			}

			//если число посередине больше левого и правого
			if numberInBytes[j-1] < numberInBytes[j] && numberInBytes[j] > numberInBytes[j+1] {
				removableIndex = j - 1
				break
			}

			if numberInBytes[removableIndex] > numberInBytes[j] {
				removableIndex = j
				//break
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
