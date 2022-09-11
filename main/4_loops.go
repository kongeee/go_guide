package main

import "fmt"

func main() {

	//there is only one looping construct -> for loop
	//in for loop, braces are not required but curly braces are required
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("for loop sum:", sum)

	//while loop is created with for keyword in go
	sum_while := 0
	fmt.Print("while loop: ")
	for sum_while < 10 {
		sum_while++
		fmt.Printf("%d%s", sum_while, " ")
	}

	//infinite loop
	for {
		//codes
	}

}
