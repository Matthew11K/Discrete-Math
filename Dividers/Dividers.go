package main

import (
	"fmt"
	"sort"
)

func main() {
	var x uint64
	fmt.Scanf("%d", &x)

	divisors := findDivisors(x)
	printGraph(x, divisors)
}

func findDivisors(x uint64) []uint64 {
	divisorsSet := make(map[uint64]struct{})
	var divisors []uint64

	for i := uint64(1); i*i <= x; i++ {
		if x%i == 0 {
			divisorsSet[i] = struct{}{}
			divisorsSet[x/i] = struct{}{}
		}
	}

	for d := range divisorsSet {
		divisors = append(divisors, d)
	}

	sort.Slice(divisors, func(i, j int) bool { return divisors[i] > divisors[j] })
	return divisors
}

func printGraph(x uint64, divisors []uint64) {
	fmt.Println("graph {")
	for _, d := range divisors {
		fmt.Printf("    %d\n", d)
	}

	for i := 0; i < len(divisors); i++ {
		u := divisors[i]
		for j := i + 1; j < len(divisors); j++ {
			v := divisors[j]
			if u%v == 0 {
				intermediate := false
				for k := 0; k < len(divisors); k++ {
					w := divisors[k]
					if w != u && w != v && u%w == 0 && w%v == 0 {
						intermediate = true
						break
					}
				}
				if !intermediate {
					fmt.Printf("    %d--%d\n", u, v)
				}
			}
		}
	}
	fmt.Println("}")
}
