package main

import "fmt"

func main() {

	//all defer items run and are pushed to stack. after function execution is done they are popped from stack. first in last out
	defer fmt.Println("first defer")
	defer fmt.Println("second defer")

	fmt.Println("method statement")
}
