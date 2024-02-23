package piscine

import "fmt"

func QuadD(x, y int) {

	for r := 1; r <= y; r++ {
		for c := 1; c <= x; c++ {
			if r == 1 && c == 1 {
				fmt.Print("A")
			} else if r == 1 && c == x {
				fmt.Print("C")
			} else if r == y && c == 1 {
				fmt.Print("A")

			} else if r == y && c == x {

				fmt.Print("C")
			} else if (r == 1 || r == y) && (c != 1 || c != x) {
				fmt.Print("B")
			} else if (c == 1 || c == x) && (r != 1 || r != y) {
				fmt.Print("B")
			} else {
				fmt.Print(" ")
			}

		}
		fmt.Println()

	}
}
