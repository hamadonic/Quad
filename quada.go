package piscine

import "fmt"

func QuadA(x, y int) {

	for r := 1; r <= y; r++ {
		for c := 1; c <= x; c++ {
			if (r == 1 && c == 1) || (r == 1 && c == x) || (r == y && c == 1) || (r == y && c == x) {
				fmt.Print("o")
			} else if (r == 1 || r == y) && (c != 1 || c != x) {
				fmt.Print("-")
			} else if (c == 1 || c == x) && (r != 1 || r != y) {
				fmt.Print("|")
			} else {
				fmt.Print(" ")
			}

		}
		fmt.Println()

	}
}
