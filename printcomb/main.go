package main

import "fmt"

func printcomb() {
	isFirst := true
	for first := 0; first <= 7; first++ {
		for second := first + 1; second <= 8; second++ {
			for third := second + 1; third <= 9; third++ {
				if !isFirst {
					fmt.Print(", ")
				}
				fmt.Printf("%d%d%d", first, second, third)
				 isFirst = false
			}
		}
	}
	fmt.Println()
}

func main() {
	printcomb()
}
