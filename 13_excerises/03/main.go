package main

import "fmt"

//Added a struct to see how that worked, worked fine
type person struct {
	Name string
	Age  int
}

func main() {
	fmt.Print("Please enter your name: ")
	var p1 = person{}
	fmt.Scan(&p1.Name)
	fmt.Println("hello", p1.Name)
}
