package main

import "fmt"

func main() {
	a := 0
	b := 0
	fmt.Println("Enter a number:")
	fmt.Scan(&a)
	fmt.Println("Enter another number:")
	fmt.Scan(&b)

	if a < b {
		fmt.Println(b, "%", a, "=", b%a)
	} else {
		fmt.Println(a, "%", b, "=", a%b)
	}
}
