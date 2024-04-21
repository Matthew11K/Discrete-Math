package main

import (
	"container/heap"
	"fmt"
)

type Item struct {
	value    [2]int
	priority int
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
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
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func dijkstra(grid [][]int, N int) int {
	dx := []int{1, 0, -1, 0}
	dy := []int{0, 1, 0, -1}

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	dist := make([][]int, N)
	for i := range dist {
		dist[i] = make([]int, N)
		for j := range dist[i] {
			dist[i][j] = 1<<31 - 1
		}
	}

	dist[0][0] = grid[0][0]
	heap.Push(&pq, &Item{[2]int{0, 0}, dist[0][0], 0})

	for pq.Len() > 0 {
		cur := heap.Pop(&pq).(*Item)
		x := cur.value[0]
		y := cur.value[1]

		if x == N-1 && y == N-1 {
			return dist[x][y]
		}

		for i := 0; i < 4; i++ {
			nx, ny := x+dx[i], y+dy[i]
			if nx >= 0 && nx < N && ny >= 0 && ny < N {
				cost := dist[x][y] + grid[nx][ny]
				if cost < dist[nx][ny] {
					dist[nx][ny] = cost
					heap.Push(&pq, &Item{[2]int{nx, ny}, cost, 0})
				}
			}
		}
	}

	return -1
}

func main() {
	var N int
	fmt.Scanf("%d", &N)

	grid := make([][]int, N)
	for i := 0; i < N; i++ {
		grid[i] = make([]int, N)
		for j := 0; j < N; j++ {
			fmt.Scanf("%d", &grid[i][j])
		}
	}

	result := dijkstra(grid, N)
	fmt.Println(result)
}
