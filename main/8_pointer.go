package main

import "fmt"

func main() {

	//pointers can hold memory address of variables. like in c.

	i := 42

	p := &i
	fmt.Printf("%v -> value of pointer| %v -> deferencing of pointer | %v -> value of the variable", p, *p, i)

	//but unlike c ho has no pointer arithmetic
	//fmt.Println(p++) -> not allowed

}
