package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//storedStr := []string{"Zero", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine"}

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	fmt.Print(input)
}
