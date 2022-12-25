package main

func rightShifted(list []int, shift int) []int {
	shiftedList := make([]int, len(list))
	for i := range list {
		shiftedList[i] = list[(i+len(list)-shift)%len(list)]
	}
	return shiftedList
}

func shiftedRow(m []int, s int) []int {
	n := make([]int, len(m))
	for i := range m {
		n[i] = m[(s+i)%len(m)]
	}
	return n
}

func shiftRows(M System, f, g Permutation) {
	for r := range M {
		switch {
		case r < 8:
			M[r] = shiftedRow(M[r], f[r])
		default:
			M[r] = shiftedRow(M[r], g[r%8])
		}
	}
}
func shiftCols(N System, f, g Permutation) {
	M := transpose(N)
	for r := range M {
		switch {
		case r < 8:
			M[r] = shiftedRow(M[r], f[r])
		default:
			M[r] = shiftedRow(M[r], g[r%8])
		}
	}
	T := transpose(M)
	for i := range N {
		N[i] = T[i]
	}
}

func shiftRowsExcept(M System, f, g Permutation, e int) {
	for r := range M {
		switch {
		case r < 8 && r != e:
			M[r] = shiftedRow(M[r], f[r])
		case r >= 8 && r != e:
			M[r] = shiftedRow(M[r], g[r%8])
		default:
		}
	}
}
func shiftColsExcept(N System, f, g Permutation, e int) {
	M := transpose(N)
	for r := range M {
		switch {
		case r < 8 && r != e:
			M[r] = shiftedRow(M[r], f[r])
		case r >= 8 && r != e:
			M[r] = shiftedRow(M[r], g[r%8])
		default:
		}
	}
	T := transpose(M)
	for i := range N {
		N[i] = T[i]
	}
}

func transpose(M System) System {
	N := zeroSystem(len(M[0]), len(M))
	for r := range M {
		for c := range M[0] {
			N[c][r] = M[r][c]
		}
	}
	return N
}

// shiftCol(m []int, s int) []int {

// }
