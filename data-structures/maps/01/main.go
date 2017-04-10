package main

import "fmt"

func main() {
	var myGreeting = make(map[string]string)
	myGreeting["Tim"] = "Good Morging"
	myGreeting["Jenny"] = "Bonjour"

	fmt.Println(myGreeting)
}
