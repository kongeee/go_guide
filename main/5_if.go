package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("results:", sqrt(2), sqrt(-5))

	fmt.Println("if with init statement:", pow(3, 2, 7)) //it is going to return 7 because of if statement in the function. (there is a limitation)
}

func sqrt(x float64) string {

	if x < 0 {
		return sqrt(-x) + "i"
	}

	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {

	//like for, the if statement can include init part. like below code a variable can be created and it is used only if statement scope (include else scope). first value of v is assigned and then condition is executed.
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%v >= %v\n", v, lim)
	}

	return lim
}
