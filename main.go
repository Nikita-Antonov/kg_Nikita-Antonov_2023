package main

import (
	"fmt"
	"os"
)

func main() {
	storedStr := []string{"Zero", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine"}

	n := 1
	for n < len(os.Args) { //Sets the max value of the loop to the length of the input
		inputValue := os.Args[n] //Setting the inputs to a variable
		n++
		for x := 0; x < len(inputValue); x++ {
			serchInt := ((inputValue[x]) - 48) //takes the ascii value of ghe inputed char, and normalizing it to an int
			fmt.Print(storedStr[serchInt])     //Uses the int to get the word, coresponding with the input number
		}
		fmt.Print(",")
	}
}
