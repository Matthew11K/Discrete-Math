package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func bfs(start int, adj [][]int, n int) []int {
	distance := make([]int, n)
	for i := range distance {
		distance[i] = -1
	}
	distance[start] = 0
	queue := []int{start}
	head := 0

	for head < len(queue) {
		current := queue[head]
		head++
		for _, neighbor := range adj[current] {
			if distance[neighbor] == -1 {
				queue = append(queue, neighbor)
				distance[neighbor] = distance[current] + 1
			}
		}
	}
	return distance
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())

	adj := make([][]int, n)
	for i := 0; i < m; i++ {
		scanner.Scan()
		edge := strings.Split(scanner.Text(), " ")
		u, _ := strconv.Atoi(edge[0])
		v, _ := strconv.Atoi(edge[1])
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	scanner.Scan()
	k, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	pivotVertices := strings.Split(scanner.Text(), " ")
	pivots := make([]int, k)
	for i, pv := range pivotVertices {
		pivots[i], _ = strconv.Atoi(pv)
	}

	allDistances := make([][]int, k)
	for i, pivot := range pivots {
		allDistances[i] = bfs(pivot, adj, n)
	}

	equidistant := make([]int, 0)
	for v := 0; v < n; v++ {
		sameDistance := true
		for i := 1; i < k; i++ {
			if allDistances[i][v] != allDistances[0][v] {
				sameDistance = false
				break
			}
		}
		if sameDistance && allDistances[0][v] != -1 {
			equidistant = append(equidistant, v)
		}
	}

	if len(equidistant) == 0 {
		fmt.Println("-")
	} else {
		for _, v := range equidistant {
			fmt.Printf("%d ", v)
		}
		fmt.Println()
	}
}
