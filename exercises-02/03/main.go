package main

import "fmt"

func greatest(numbers ...int) int {
	var largest int
	for _, i := range numbers {
		if i > largest {
			largest = i
		}
	}
	return largest
}

func main() {
	greatest := greatest(24, 51, 123, 255, 1098, 240, 1)
	fmt.Println(greatest)
}
