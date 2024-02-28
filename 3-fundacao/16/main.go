package main

import (
	"fmt"
)

func main() {
	var minhaVar interface{} = "Wesley Willians"
	println(minhaVar.(string))
	res, ok := minhaVar.(int) // res will be 0 and ok will be false. Good for type assertions.
	fmt.Printf("O valor de res é %v e o resultado de ok é %v\n", res, ok)
	res2 := minhaVar.(int) // it will panic, since we are trying to convert a string to an int and not getting the ok value.
	fmt.Printf("O valor de res2 é %v\n", res2)
}
