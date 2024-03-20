package main

import (
	"fmt"
)

func myPanic1() {
	panic("panic1")
}

func myPanic2() {
	panic("panic2")
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			if r == "panic1" {
				fmt.Println("Recovered in main:", r)
			}
			if r == "panic2" {
				fmt.Println("Recovered in main:", r)
			}
		}
	}()

	myPanic1()
	myPanic2() // won't be executed
}
