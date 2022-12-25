package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	test_encode(24)
	//test()
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
	var f, g, h, h_inv Permutation
	rand.Seed(time.Now().UnixNano())
	H := make(map[string]bool)

	S := randomSystem(16, 16)
	I := 1000
	for i := 0; i < I; i++ {
		f, _ = permFrom(octalFrom(flatten(S, false)))
		g, _ = permFrom(octalFrom(flatten(S, true)))
		h = composition(f, g)
		h_inv = inverse(h)
		fmt.Printf("h     = %v\n", h)
		fmt.Printf("h_inv = %v\n", h_inv)

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

func test_encode(n int) {
	rand.Seed(time.Now().UnixNano())
	S := randomSystem(16, 16)
	S.fancy()

	for i := 0; i < 30; i++ {
		plain := randomPlain(n)
		code := S.encode(plain)
		fmt.Printf("\nf%v = %v\n", plain, code)
		decode := S.decode(code)
		fmt.Printf("\ng%v = %v\n", code, decode)

	}

}
