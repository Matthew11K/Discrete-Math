package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var (
	scanner           = bufio.NewScanner(os.Stdin)
	adjList           map[int][]int
	index, sccCounter int
	indices, lowLink  []int
	inStack           []bool
	stack             []int
	scc               []int
	sccAdjList        map[int][]int
)

func main() {
	scanner.Split(bufio.ScanWords)
	var N, M int
	scan(&N)
	scan(&M)
	adjList = make(map[int][]int)
	for i := 0; i < M; i++ {
		var u, v int
		scan(&u)
		scan(&v)
		adjList[u] = append(adjList[u], v)
	}

	indices = make([]int, N)
	lowLink = make([]int, N)
	inStack = make([]bool, N)
	scc = make([]int, N)
	for i := range indices {
		indices[i] = -1
	}

	for i := 0; i < N; i++ {
		if indices[i] == -1 {
			strongConnect(i)
		}
	}

	sccAdjList = make(map[int][]int)
	for u, neighbors := range adjList {
		for _, v := range neighbors {
			sccU, sccV := scc[u], scc[v]
			if sccU != sccV {
				sccAdjList[sccU] = append(sccAdjList[sccU], sccV)
			}
		}
	}

	inDegree := make([]int, sccCounter)
	for _, neighbors := range sccAdjList {
		for _, v := range neighbors {
			inDegree[v]++
		}
	}

	var base []int
	for i := 0; i < sccCounter; i++ {
		if inDegree[i] == 0 {
			base = append(base, findMinVertexInSCC(i, N))
		}
	}

	sort.Ints(base)
	for _, v := range base {
		fmt.Printf("%d ", v)
	}
}

func strongConnect(v int) {
	indices[v] = index
	lowLink[v] = index
	index++
	stack = append(stack, v)
	inStack[v] = true

	for _, w := range adjList[v] {
		if indices[w] == -1 {
			strongConnect(w)
			lowLink[v] = min(lowLink[v], lowLink[w])
		} else if inStack[w] {
			lowLink[v] = min(lowLink[v], indices[w])
		}
	}

	if lowLink[v] == indices[v] {
		var w int
		for {
			w, stack = stack[len(stack)-1], stack[:len(stack)-1]
			inStack[w] = false
			scc[w] = sccCounter
			if w == v {
				break
			}
		}
		sccCounter++
	}
}

func findMinVertexInSCC(sccNum, N int) int {
	minVertex := N
	for v := 0; v < N; v++ {
		if scc[v] == sccNum && v < minVertex {
			minVertex = v
		}
	}
	return minVertex
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func scan(a *int) {
	scanner.Scan()
	fmt.Sscan(scanner.Text(), a)
}
