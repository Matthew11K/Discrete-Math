package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Edge struct {
	to, weight int
}

type PriorityQueue []*Item

type Item struct {
	vertex int
	weight int
	index  int
}

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].weight < pq[j].weight
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) update(item *Item, vertex int, weight int) {
	item.vertex = vertex
	item.weight = weight
	heap.Fix(pq, item.index)
}

func readGraph() (int, [][]Edge) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	M, _ := strconv.Atoi(scanner.Text())

	graph := make([][]Edge, N)
	for i := 0; i < M; i++ {
		scanner.Scan()
		parts := strings.Split(scanner.Text(), " ")
		u, _ := strconv.Atoi(parts[0])
		v, _ := strconv.Atoi(parts[1])
		length, _ := strconv.Atoi(parts[2])

		graph[u] = append(graph[u], Edge{v, length})
		graph[v] = append(graph[v], Edge{u, length})
	}

	return N, graph
}

func prim(N int, graph [][]Edge) int {
	totalWeight := 0
	visited := make([]bool, N)
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	heap.Push(&pq, &Item{vertex: 0, weight: 0})

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		v := item.vertex

		if visited[v] {
			continue
		}

		visited[v] = true
		totalWeight += item.weight

		for _, edge := range graph[v] {
			if !visited[edge.to] {
				heap.Push(&pq, &Item{vertex: edge.to, weight: edge.weight})
			}
		}
	}

	return totalWeight
}

func main() {
	N, graph := readGraph()
	result := prim(N, graph)
	fmt.Println(result)
}
