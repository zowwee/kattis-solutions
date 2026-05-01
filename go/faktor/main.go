package main

import "fmt"

func main() {
	var x, y int
	fmt.Scanln(&x, &y)

	// articlesToPublish := x
	// requiredImpactFactor := y
	minCitations := x * (y - 1)
	maxiCitations := x * y

	for i := minCitations; i < maxiCitations+1; i++ {
		if float64(i)/float64(x) > float64(y)-1 {
			fmt.Println(i)
			break
		}
	}
}
