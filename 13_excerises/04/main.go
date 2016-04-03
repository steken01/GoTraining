package main

import "fmt"

func main() {
	var smallNumber int
	var largeNumber int
	fmt.Print("Enter a small number: ")
	fmt.Scan(&smallNumber)
	fmt.Print("Enter a large number: ")
	fmt.Scan(&largeNumber)
	fmt.Print("The large number remainder when divided by the small number is: ", largeNumber%smallNumber)
}
