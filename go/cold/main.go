package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)

	sum := 0

	for i := 0; i < n; i++ {
		var x int
		fmt.Scan(&x)

		if x < 0 {
			sum++
		}
	}
	fmt.Println(sum)
}
