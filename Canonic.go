package main

import (
	"fmt"
)

func main() {
	var n, m, q0 int

	fmt.Scan(&n)
	fmt.Scan(&m)
	fmt.Scan(&q0)

	transition := make([][]int, n)
	for i := range transition {
		transition[i] = make([]int, m)
		for j := range transition[i] {
			fmt.Scan(&transition[i][j])
		}
	}

	output := make([][]string, n)
	for i := range output {
		output[i] = make([]string, m)
		for j := range output[i] {
			fmt.Scan(&output[i][j])
		}
	}

	visited := make([]bool, n)
	order := make([]int, n)
	label := 0
	var dfs func(int)
	dfs = func(v int) {
		visited[v] = true
		order[v] = label
		label++

		for input := 0; input < m; input++ {
			next := transition[v][input]
			if !visited[next] {
				dfs(next)
			}
		}
	}

	dfs(q0)

	newTransition := make([][]int, n)
	newOutput := make([][]string, n)
	for i := 0; i < n; i++ {
		newTransition[order[i]] = make([]int, m)
		newOutput[order[i]] = make([]string, m)
		for j := 0; j < m; j++ {
			newTransition[order[i]][j] = order[transition[i][j]]
			newOutput[order[i]][j] = output[i][j]
		}
	}

	fmt.Println(n)
	fmt.Println(m)
	fmt.Println(order[q0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Printf("%d ", newTransition[i][j])
		}
		fmt.Println()
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Printf("%s ", newOutput[i][j])
		}
		fmt.Println()
	}
}
