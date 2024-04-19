package main

import (
	"fmt"
	"math"
	"sort"
)

type Point struct {
	X, Y float64
}

type Edge struct {
	Start, End int
	Weight     float64
}

type UnionFind struct {
	Parent, Rank []int
}

func NewUnionFind(size int) *UnionFind {
	uf := &UnionFind{
		Parent: make([]int, size),
		Rank:   make([]int, size),
	}
	for i := range uf.Parent {
		uf.Parent[i] = i
	}
	return uf
}

func (uf *UnionFind) Find(x int) int {
	if uf.Parent[x] != x {
		uf.Parent[x] = uf.Find(uf.Parent[x])
	}
	return uf.Parent[x]
}

func (uf *UnionFind) Union(x, y int) {
	rootX := uf.Find(x)
	rootY := uf.Find(y)
	if rootX != rootY {
		if uf.Rank[rootX] > uf.Rank[rootY] {
			uf.Parent[rootY] = rootX
		} else if uf.Rank[rootX] < uf.Rank[rootY] {
			uf.Parent[rootX] = rootY
		} else {
			uf.Parent[rootY] = rootX
			uf.Rank[rootX]++
		}
	}
}

func Kruskal(edges []Edge, n int) float64 {
	uf := NewUnionFind(n)
	var totalWeight float64
	for _, e := range edges {
		if uf.Find(e.Start) != uf.Find(e.End) {
			uf.Union(e.Start, e.End)
			totalWeight += e.Weight
		}
	}
	return totalWeight
}

func distance(p1, p2 Point) float64 {
	dx := p1.X - p2.X
	dy := p1.Y - p2.Y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	var n int
	fmt.Scanf("%d", &n)

	points := make([]Point, n)
	for i := range points {
		fmt.Scanf("%f %f", &points[i].X, &points[i].Y)
	}

	var edges []Edge
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			dist := distance(points[i], points[j])
			edges = append(edges, Edge{i, j, dist})
		}
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].Weight < edges[j].Weight
	})

	minTotalLength := Kruskal(edges, n)
	fmt.Printf("%.2f\n", minTotalLength)
}
