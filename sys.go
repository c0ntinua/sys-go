package main

import (
	"fmt"
	"math/rand"
)

type System [][]int

func zeroSystem(r, c int) System {
	M := make([][]int, r)
	for i := range M {
		M[i] = make([]int, c)
	}
	return M
}

func randomSystem(r, c int) System {
	M := make(System, r)
	for i := range M {
		M[i] = make([]int, c)
	}
	M.randomize()
	return M
}

func (M System) randomize() {
	for i := range M {
		for j := range M[i] {
			M[i][j] = rand.Intn(2)
		}
		fmt.Println()
	}
}

func flatten(M System, reverse bool) []int {
	join := make([]int, len(M)*len(M[0]))
	for i := range M {
		for j := range M[i] {
			switch {
			case reverse:
				join[len(join)-1-len(M[0])*i-j] = M[i][j]
			default:
				join[len(M[0])*i+j] = M[i][j]
			}
		}
	}
	return join
}
