package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	test()
}

func trial() {
	// rand.Seed(time.Now().UnixNano())
	// S := randomSystem(16, 16)
	// S.fancy()
	// f, _ := permFrom(octalFrom(flatten(S, false)))
	// g, _ := permFrom(octalFrom(flatten(S, true)))
	// h := composition(f, g)
	// fmt.Printf("f    = %v\n", f)
	// fmt.Printf("g    = %v\n", g)
	// fmt.Printf("h    = %v\n", h)
	// //shiftCols(S, g, f)
	// shiftRows(S, f, g)
	// S.fancy()
	// f, _ = permFrom(octalFrom(flatten(S, false)))
	// g, _ = permFrom(octalFrom(flatten(S, true)))
	// h = composition(f, g)
	// fmt.Printf("f    = %v\n", f)
	// fmt.Printf("g    = %v\n", g)
	// fmt.Printf("h    = %v\n", h)
	// shiftCols(S, f, g)
	// S.fancy()

}
func test() {
	var f, g, h Permutation
	rand.Seed(time.Now().UnixNano())
	H := make(map[string]bool)

	S := randomSystem(16, 16)
	I := 1000
	for i := 0; i < I; i++ {
		f, _ = permFrom(octalFrom(flatten(S, false)))
		g, _ = permFrom(octalFrom(flatten(S, true)))
		h = composition(f, g)
		shiftRows(S, f, g)
		shiftCols(S, g, f)
		_, result := H[asString(h)]
		switch {
		case result:
			fmt.Printf("%d repeat as %v \n", i, h)
		}
		H[asString(h)] = true
	}
	fmt.Printf("Found %v out of %v unique permutations...n", len(H), I)

}
