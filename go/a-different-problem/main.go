package main

import "fmt"

func main() {
	for {
		var x, y int

		// check if input ends
		_, err := fmt.Scanln(&x, &y)
		if err != nil {
			break
		}

		difference := x - y

		if difference < 0 {
			difference = -difference
		}

		fmt.Println(difference)
	}
}
