package main

func main() {
	// Memória -> Endreço -> Valor
	a := 10
	var ponteiro *int = &a // & access the address of a variable
	*ponteiro = 20 // * access the value of the address
	b := &a
	*b = 30
	println(a)
}
