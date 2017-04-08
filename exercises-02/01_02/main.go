package main

import "fmt"

func half(x int) (float64, bool) {
	return float64(x) / 2, x%2 == 0
}

func main() {
	h, isEven := half(1)
	fmt.Println(h, isEven)
}
