package main

import (
	"fmt"
	"sort"
)

func main() {
	i := []int{8, 10, 6, 3}

	fmt.Println(i)
	sort.Ints(i)
	fmt.Println(i)
}
