package main

import "fmt"

func main() {
	var input string
	fmt.Scanln(&input)

	numberOfASCII := 0
	for _, character := range input {
		if (character >= 65 && character <= 90) || (character >= 97 && character <= 122) {
			numberOfASCII += 1
		}
	}

	fmt.Println(numberOfASCII)
}
