package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

type Edge struct {
	to, color int
}

type State struct {
	node, length int
	path         []int
}

type PriorityQueue []State

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	if pq[i].length != pq[j].length {
		return pq[i].length < pq[j].length
	}
	for k := 0; k < len(pq[i].path) && k < len(pq[j].path); k++ {
		if pq[i].path[k] != pq[j].path[k] {
			return pq[i].path[k] < pq[j].path[k]
		}
	}
	return len(pq[i].path) < len(pq[j].path)
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(State))
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())

	adjList := make(map[int][]Edge)
	for i := 0; i < m; i++ {
		scanner.Scan()
		a, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		b, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		c, _ := strconv.Atoi(scanner.Text())

		adjList[a] = append(adjList[a], Edge{to: b, color: c})
		adjList[b] = append(adjList[b], Edge{to: a, color: c})
	}

	path, length := dijkstra(1, n, adjList)
	fmt.Println(length)
	for _, color := range path {
		fmt.Printf("%d ", color)
	}
}

func dijkstra(start, end int, adjList map[int][]Edge) ([]int, int) {
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, State{node: start, length: 0, path: []int{}})

	visited := make(map[int][]int)

	for pq.Len() > 0 {
		current := heap.Pop(&pq).(State)
		if current.node == end {
			return current.path, current.length
		}
		if path, found := visited[current.node]; found && len(path) <= len(current.path) && lexicographicalLess(path, current.path) {
			continue
		}
		visited[current.node] = current.path

		for _, edge := range adjList[current.node] {
			newPath := make([]int, len(current.path)+1)
			copy(newPath, current.path)
			newPath[len(newPath)-1] = edge.color

			heap.Push(&pq, State{node: edge.to, length: current.length + 1, path: newPath})
		}
	}
	return nil, -1
}

func lexicographicalLess(a, b []int) bool {
	for i := 0; i < len(a) && i < len(b); i++ {
		if a[i] != b[i] {
			return a[i] < b[i]
		}
	}
	return len(a) < len(b)
}
