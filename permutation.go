package main

import (
	"strconv"
	"strings"
)

type Permutation []int

// func (s System) f(bits int) Permutation {

// }
func octalFrom(binaryList []int) []int {
	octalList := make([]int, len(binaryList)/3)
	for i := 0; i < len(octalList); i++ {
		octalList[i] = 4*binaryList[3*i] + 2*binaryList[3*i+1] + binaryList[3*i+2]
	}
	return octalList
}

func permFrom(octalList []int) ([]int, bool) {
	have := make(map[int]bool, 8)
	f := make([]int, 8)
	counter := 0
	for _, v := range octalList {
		if !have[v] {
			have[v] = true
			f[counter] = v
			counter++
		}
	}
	if counter == 8 {
		return f, true
	}
	return f, false
}

func composition(f, g Permutation) Permutation {
	h := make(Permutation, 8)
	for i := range h {
		h[i] = f[g[i]]
	}
	return h
}
func asString(p Permutation) string {
	var s []string
	for i := range p {
		s = append(s, strconv.Itoa(p[i]))
	}
	return strings.Join(s, "")
}
func inverse(f Permutation) Permutation {
	g := make(Permutation, len(f))
	j := 0
	for i := range f {
		for j = 0; j != f[i]; j++ {
		}
		g[j] = i
	}
	return g
}
