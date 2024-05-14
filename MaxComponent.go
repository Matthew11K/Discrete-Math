package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Graph struct {
	adjacencyList [][]int
	edges         [][]int
}

func newGraph(n int) *Graph {
	return &Graph{
		adjacencyList: make([][]int, n),
		edges:         [][]int{},
	}
}

func (g *Graph) addEdge(u, v int) {
	g.adjacencyList[u] = append(g.adjacencyList[u], v)
	g.adjacencyList[v] = append(g.adjacencyList[v], u)
	g.edges = append(g.edges, []int{u, v})
}

func (g *Graph) findLargestComponent() (component []int, edgeCount int) {
	n := len(g.adjacencyList)
	visited := make([]bool, n)
	var largestComponent []int
	maxEdges := -1

	var dfs func(v int, comp *[]int) int
	dfs = func(v int, comp *[]int) int {
		visited[v] = true
		*comp = append(*comp, v)
		edges := 0

		for _, u := range g.adjacencyList[v] {
			edges++
			if !visited[u] {
				edges += dfs(u, comp)
			}
		}

		return edges
	}

	for i := 0; i < n; i++ {
		if !visited[i] {
			var comp []int
			edges := dfs(i, &comp) / 2
			if len(comp) > len(largestComponent) ||
				(len(comp) == len(largestComponent) && edges > maxEdges) ||
				(len(comp) == len(largestComponent) && edges == maxEdges && comp[0] < largestComponent[0]) {
				largestComponent = comp
				maxEdges = edges
			}
		}
	}

	return largestComponent, maxEdges
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())

	graph := newGraph(n)

	for i := 0; i < m; i++ {
		scanner.Scan()
		parts := strings.Split(scanner.Text(), " ")
		u, _ := strconv.Atoi(parts[0])
		v, _ := strconv.Atoi(parts[1])
		graph.addEdge(u, v)
	}

	largestComponent, _ := graph.findLargestComponent()

	isInLargestComponent := make(map[int]bool)
	for _, v := range largestComponent {
		isInLargestComponent[v] = true
	}

	fmt.Println("graph {")
	for i := 0; i < n; i++ {
		if isInLargestComponent[i] {
			fmt.Printf("  %d [color=red];\n", i)
		} else {
			fmt.Printf("  %d;\n", i)
		}
	}

	for _, edge := range graph.edges {
		u, v := edge[0], edge[1]
		if isInLargestComponent[u] && isInLargestComponent[v] {
			fmt.Printf("  %d -- %d [color=red];\n", u, v)
		} else {
			fmt.Printf("  %d -- %d;\n", u, v)
		}
	}

	fmt.Println("}")
}
