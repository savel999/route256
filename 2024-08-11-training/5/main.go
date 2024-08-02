package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
)

var sanitizeJsonReg = regexp.MustCompile(`(?m)("[^"]+":)*(\[\]|\{\})`)
var sanitizeMultipleCommas = regexp.MustCompile(`(?m),{2,}`)
var sanitizeLeadCommas = regexp.MustCompile(`(?m)([\{\[]),`)
var sanitizeTrailingCommas = regexp.MustCompile(`(?m),([\}\]])`)

func sanitizeJSONPart(src []byte) []byte {
	dst := &bytes.Buffer{}
	if err := json.Compact(dst, src); err != nil {
		panic(err)
	}

	src = dst.Bytes()

	for {
		replaced := sanitizeJsonReg.ReplaceAll(src, nil)
		replaced = sanitizeMultipleCommas.ReplaceAll(replaced, []byte(","))
		replaced = sanitizeLeadCommas.ReplaceAll(replaced, []byte("$1"))
		replaced = sanitizeTrailingCommas.ReplaceAll(replaced, []byte("$1"))

		if len(src) == len(replaced) {
			return replaced
		}

		src = replaced
	}

}

func mergeJSONParts(parts ...[]byte) []byte {
	result := []byte("[")
	result = append(result, bytes.Join(parts, []byte(","))...)
	result = append(result, []byte("]")...)

	return result
}

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var cnt int
	fmt.Fscan(in, &cnt)

	sanitizedParts := make([][]byte, cnt)

	for i := 0; i < cnt; i++ {
		var linesCnt int

		fmt.Fscan(in, &linesCnt)

		var jsonContent []byte

		for j := 0; j < linesCnt+1; j++ {
			line, _, _ := in.ReadLine()

			jsonContent = append(jsonContent, line...)
		}

		sanitizedParts[i] = sanitizeJSONPart(jsonContent)
	}

	fmt.Fprintln(out, string(mergeJSONParts(sanitizedParts...)))
}
