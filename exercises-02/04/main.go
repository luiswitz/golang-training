package main

import "fmt"

func main() {
	expression := (true && false) || (false && true) || !(false && false)
	fmt.Println(expression)
}
