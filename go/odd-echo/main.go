package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)

	for i := 1; i < n+1; i++ {
		var x string
		fmt.Scanln(&x)

		if i%2 != 0 {
			fmt.Println(x)
		}
	}
}
