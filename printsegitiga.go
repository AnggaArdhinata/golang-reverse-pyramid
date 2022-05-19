package main

import (
	"fmt"
)

func main() {

	printSegitiga(8)
}

func printSegitiga(number int) {

	if number <= 1 {
		fmt.Println("Data harus lebih dari satu")

	} else {

		for x := number; x >= 1; x-- {
			for y := number; y >= x; y-- {
				fmt.Print(" ")
			}
			for z := 1; z <= x; z++ {
				fmt.Print("* ")
			}
			fmt.Println()
		}
	}

}
