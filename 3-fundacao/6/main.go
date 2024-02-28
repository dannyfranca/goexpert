package main

import "fmt"

func main() {
	s := []int{10, 20, 30, 50, 60, 70, 80, 90, 100}
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s) // len=9 cap=16 [10 20 30 50 60 70 80 90 100]

	fmt.Printf("len=%d cap=%d %v\n", len(s[:0]), cap(s[:0]), s[:0]) // len=0 cap=9 [] - a slice with zero length

	fmt.Printf("len=%d cap=%d %v\n", len(s[:4]), cap(s[:4]), s[:4]) // len=4 cap=9 [10 20 30 50] - a slice with the first 4 elements

	fmt.Printf("len=%d cap=%d %v\n", len(s[2:]), cap(s[2:]), s[2:]) // len=7 cap=7 [30 50 60 70 80 90] - a slice with the elements from index 2 to the end

	s = append(s, 110) // append adds a new element to the end of the slice, increasing its length by 1 and current capacity is doubled
	fmt.Printf("len=%d cap=%d %v\n", len(s[:2]), cap(s[:2]), s[:2]) // len=2 cap=18 [10 20] - a slice with the first 2 elements
}
