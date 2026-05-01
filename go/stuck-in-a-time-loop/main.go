package main

import "fmt"

func main() {
	var N int

	fmt.Scanln(&N)

	for i := 1; i < N+1; i++ {
		fmt.Println(i, "", "Abracadabra")
	}
}
