package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scanln(&n)

	reader := bufio.NewReader(os.Stdin)
	numbers, _ := reader.ReadString('\n')
	sum := 0
	numbersSlice := strings.Fields(numbers)

	for i, number := range numbersSlice {
		if i < n {
			number = string(number)
			number, _ = strconv.Atoi(number)
			sum += number
		}
		fmt.Println(sum)
	}
}
