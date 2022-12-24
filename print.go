package main

import "fmt"

func (M System) display() {
	for i := range M {
		for j := range M[i] {
			fmt.Printf("%2d", M[i][j])
		}
		fmt.Println()
	}
}
func (M System) fancy() {
	for i := range M {
		for j := range M[i] {
			switch {
			case M[i][j] == 0:
				fmt.Printf("O ")
			default:
				fmt.Printf("| ")
			}
		}
		fmt.Println()
	}
}
