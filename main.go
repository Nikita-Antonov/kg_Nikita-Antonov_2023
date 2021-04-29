package main

import (
	"fmt"
	"os"
)

func main() {
	storedStr := []string{"Zero", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine"}

	n := 1
	for n < len(os.Args) {
		inputValue := os.Args[n]
		var numInInput int = len(inputValue)
		n++
		for x := 0; x < numInInput; x++ {
			serchInt := ((inputValue[x]) - 48)
			fmt.Print(storedStr[serchInt])
		}
		fmt.Print(",")
	}

}
