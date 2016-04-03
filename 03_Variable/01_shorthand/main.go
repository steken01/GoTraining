package main

import "fmt"

func main() {

	a := 2
	b := "two"
	c := false
	d := 4
	e := 4.14

	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)
	fmt.Printf("%v \n\n", e)

	fmt.Printf("%T \n", a)
	fmt.Printf("%T \n", b)
	fmt.Printf("%T \n", c)
	fmt.Printf("%T \n", d)
	fmt.Printf("%T \n", e)
}
