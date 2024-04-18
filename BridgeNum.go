package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	adj      [][]int
	visited  []bool
	disc     []int
	low      []int
	parent   []int
	bridges  int
	timeStep int
)

func dfs(u int) {
	visited[u] = true
	disc[u] = timeStep
	low[u] = timeStep
	timeStep++

	for _, v := range adj[u] {
		if !visited[v] {
			parent[v] = u
			dfs(v)

			if low[v] > disc[u] {
				bridges++
			}

			low[u] = min(low[u], low[v])
		} else if v != parent[u] {
			low[u] = min(low[u], disc[v])
		}
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())

	adj = make([][]int, n)
	visited = make([]bool, n)
	disc = make([]int, n)
	low = make([]int, n)
	parent = make([]int, n)
	for i := range parent {
		parent[i] = -1
	}

	for i := 0; i < m; i++ {
		scanner.Scan()
		edge := scanner.Text()
		var u, v int
		fmt.Sscanf(edge, "%d %d", &u, &v)
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	for i := 0; i < n; i++ {
		if !visited[i] {
			dfs(i)
		}
	}

	fmt.Println(bridges)
}
