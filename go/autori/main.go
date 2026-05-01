package main

import (
	"fmt"
	"strings"
)

func main() {
	var x string

	fmt.Scanln(&x)

	parts := strings.Split(x, "-")

	Initials := ""

	for _, part := range parts {
		a := string(part[0])
		Initials += a
	}

	fmt.Println(Initials)
}
