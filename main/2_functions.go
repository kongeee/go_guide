package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//using math library function
	fmt.Println("Random number is:", rand.Intn(10))

	//using custom function
	fmt.Println("add function result:", add(15, 8))

	//functions can return multiple variables at the same time
	a, b := swap("hello", "world")
	fmt.Println("swap function result:", a, b)

	//functions can return any number of variables without in return statement. variables that is going to be returned is written at the top of the function. at the end of the function there is only return statement. this type of return statement is called "naked" return
	fmt.Print("split function result: ")
	fmt.Println(split(17))
}

func add(x int, y int) int { //this is also rewritten like this func add(x, y int) int {.......}
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
