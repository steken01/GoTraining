package main

import "fmt"

func main() {

	for i := 0; i < 2000; i++ {
		fmt.Printf("Decimal: %d\tBinary: %b\tHexadecimal: %#X\tUnicode: %c\n", i, i, i, i)
		//Added %c to see that it was and it was the Uncode character corresponding to that number
	}
}
