package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	var a, b string
	for _, character := range input {
		a = string(character)
		if a != b {
			fmt.Print(a)
		}
		b = a
	}
}
