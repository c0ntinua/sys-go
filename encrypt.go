package main

import "math/rand"

func (S System) encode(plain []int) []int {
	T := systemCopy(S)
	var f, g, h, fg, gf Permutation
	code := make([]int, len(plain))
	for i := range plain {
		f, _ = permFrom(octalFrom(flatten(T, false)))
		g, _ = permFrom(octalFrom(flatten(T, true)))
		h = composition(f, g)
		code[i] = h[plain[i]]
		fg = composition(f, g)
		gf = composition(g, f)
		shiftRowsExcept(T, fg, gf, code[i])
		shiftColsExcept(T, gf, fg, code[i])
	}
	return code
}
func (S System) decode(code []int) []int {
	T := systemCopy(S)
	var f, g, h, fg, gf Permutation
	plain := make([]int, len(code))
	for i := range code {
		f, _ = permFrom(octalFrom(flatten(T, false)))
		g, _ = permFrom(octalFrom(flatten(T, true)))
		h = inverse(composition(f, g))
		plain[i] = h[code[i]]
		fg = composition(f, g)
		gf = composition(g, f)
		shiftRowsExcept(T, fg, gf, code[i])
		shiftColsExcept(T, gf, fg, code[i])
	}
	return plain
}

func randomPlain(n int) []int {
	plain := make([]int, n)
	for i := range plain {
		plain[i] = rand.Intn(8)
	}
	return plain
}

func systemCopy(S System) System {
	T := zeroSystem(len(S), len(S[0]))
	for r := range S {
		for c := range S[0] {
			T[r][c] = S[r][c]
		}
	}
	return T
}
