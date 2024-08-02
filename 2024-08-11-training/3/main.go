package main

import (
	"bufio"
	"fmt"
	"os"
)

type Node struct {
	Value int
	Nodes []Node
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var cnt int
	fmt.Fscan(in, &cnt)

	for i := 0; i < cnt; i++ {
		var digitsCnt int
		fmt.Fscan(in, &digitsCnt)

		digits := make([]int, digitsCnt)

		for j := 0; j < digitsCnt; j++ {
			var digitValue int
			fmt.Fscan(in, &digitValue)

			digits[j] = digitValue
		}

		fmt.Fprintln(out, getRootValue(numsToTree(digits)))
	}
}

func numsToTree(nums []int) map[int]Node {
	nodesMap := make(map[int]Node)
	j, digitsCnt := 0, len(nums)

	for j < digitsCnt {
		nodeValue, childNotesCnt := nums[j], nums[j+1]
		j += 2

		node := Node{Value: nodeValue, Nodes: make([]Node, childNotesCnt)}

		for k := 0; k < childNotesCnt; k++ {
			childNodeValue := nums[j]

			node.Nodes[k] = Node{Value: childNodeValue}

			j++
		}

		nodesMap[nodeValue] = node
	}

	return nodesMap
}

func getRootValue(nodesMap map[int]Node) int {
	if len(nodesMap) == 1 {
		for nodeVal := range nodesMap {
			return nodeVal
		}
	}

	for nodeVal, node := range nodesMap {
		if len(node.Nodes) == 0 {
			continue
		}

		var newChildNodes []Node

		for _, childNode := range node.Nodes {
			if foundedNode, ok := nodesMap[childNode.Value]; len(foundedNode.Nodes) == 0 && ok {
				delete(nodesMap, foundedNode.Value)

				continue
			}

			newChildNodes = append(newChildNodes, childNode)
		}

		nodesMap[nodeVal] = Node{Value: nodeVal, Nodes: newChildNodes}
	}

	return getRootValue(nodesMap)
}
