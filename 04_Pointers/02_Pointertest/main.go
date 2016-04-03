package main

import "fmt"

func main() {
	var a = 10
	var b = 11
	var pa = &a
	var pb = &b

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(pa)
	fmt.Println(pb)
	//	pv = pa + pb
	//did not work :( + not definde on pointer
	*pa = 100
	fmt.Println(a)
}
