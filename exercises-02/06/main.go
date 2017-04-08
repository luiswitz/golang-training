package main

import (
	"fmt"
	"time"
)

func fibonnaci_even_sum(limit int) int {
	sum, x, y := 2, 1, 2
	for y <= limit {
		x, y = y, x+y
		if y%2 == 0 {
			sum += y
		}
	}
	return sum
}

func main() {
	start := time.Now()
	var result int
	for i := 0; i < 1000000; i++ {
		result = fibonnaci_even_sum(4000000)
	}
	elapsed := time.Since(start)
	fmt.Println("Result ", result)
	fmt.Println("Took %s", elapsed)
}
