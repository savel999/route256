package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var employeesCnt int
	fmt.Fscan(in, &employeesCnt)

	employeesMap := make(map[string][][]byte, employeesCnt)

	for i := 0; i < employeesCnt; i++ {
		var login []byte
		fmt.Fscan(in, &login)

		clonedLogin := bytes.Clone(login)
		sort.Slice(clonedLogin, func(i, j int) bool { return clonedLogin[i] < clonedLogin[j] })

		clonedLineAscSymbols := string(clonedLogin)

		employeesMap[clonedLineAscSymbols] = append(employeesMap[clonedLineAscSymbols], login)
	}

	var recruitsCnt int
	fmt.Fscan(in, &recruitsCnt)

	for i := 0; i < recruitsCnt; i++ {
		var login []byte
		fmt.Fscan(in, &login)

		clonedLogin := bytes.Clone(login)
		sort.Slice(clonedLogin, func(i, j int) bool { return clonedLogin[i] < clonedLogin[j] })

		abc := hasSameLogin(login, string(clonedLogin), employeesMap)

		fmt.Fprintf(out, "%d\n", abc)
	}
}

func hasSameLogin(login []byte, ascedLoginSymbols string, loginsMap map[string][][]byte) int {
	if sameLogins, ok := loginsMap[ascedLoginSymbols]; ok {
		for _, sameLogin := range sameLogins {
			// если логины полностью совпали
			if bytes.Equal(sameLogin, login) {
				return 1
			}

			// если совпадает при условии перестанови 2х соседних символов
			clonedLogin := bytes.Clone(login)

			for i := 0; i < len(clonedLogin)-1; i++ {
				if clonedLogin[i] != sameLogin[i] {
					clonedLogin[i], clonedLogin[i+1] = clonedLogin[i+1], clonedLogin[i]

					if bytes.Equal(sameLogin, clonedLogin) {
						return 1
					}

					break
				}
			}
		}
	}

	return 0
}
