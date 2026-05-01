package main

import "fmt"

func printRange(small, large int) {
	for i := small + 1; i <= large+1; i++ {
		fmt.Println(i)
	}
}

func main() {
	var x, y int
	fmt.Scanln(&x, &y)

	//identify the smaller and bigger die
	if x < y {
		small := x
		large := y

		printRange(small, large)
	} else {
		small := y
		large := x

		printRange(small, large)
	}

}
