package main

import "fmt"

//a struct is a collection of data fields
type Coordinate struct {
	X int
	Y int
}

func main() {
	c := Coordinate{1, 98}
	cptr := &c
	fmt.Println("value of c variable (Coordinate type):", c)
	fmt.Println("X value of Coordinate struct is:", c.X)
	fmt.Println("Accessing Y value with cptr:", cptr.Y) //dont need (*cptr).Y .go permits cptr.Y so programmers dont need any deferencing.

	//value of the type can be changed
	c.Y = 5
	cptr.X = 687
	fmt.Println("New value of the coordinate:", c)

	//definitions and assigning values for structs
	var (
		v1 = Coordinate{1, 2}            //classic assignment
		v2 = Coordinate{X: 1, Y: 65}     //assignment with field name
		v3 = Coordinate{X: 98}           //Y value is 0 as default
		v4 = Coordinate{}                //X 0, Y 0
		v5 = &Coordinate{X: 656, Y: 597} //address of coordinate struct is assigned
	)
	fmt.Println("Values:", v1, v2, v3, v4, v5)

}
