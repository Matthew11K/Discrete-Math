package main

import (
	"fmt"
)

func main() {
	var n, m, q0 int

	fmt.Scanf("%d", &n)
	fmt.Scanf("%d", &m)
	fmt.Scanf("%d", &q0)

	transitionMatrix := make([][]int, n)
	for i := range transitionMatrix {
		transitionMatrix[i] = make([]int, m)
		for j := range transitionMatrix[i] {
			fmt.Scanf("%d", &transitionMatrix[i][j])
		}
	}

	outputMatrix := make([][]string, n)
	for i := range outputMatrix {
		outputMatrix[i] = make([]string, m)
		for j := range outputMatrix[i] {
			fmt.Scanf("%s", &outputMatrix[i][j])
		}
	}

	fmt.Println("digraph {")
	fmt.Println("\trankdir = LR")
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			inputSignal := rune('a' + j)
			fmt.Printf("\t%d -> %d [label = \"%c(%s)\"]\n", i, transitionMatrix[i][j], inputSignal, outputMatrix[i][j])
		}
	}
	fmt.Println("}")
}
