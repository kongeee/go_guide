package main

import (
	"fmt"
	"math/cmplx"
)

var k, l, m bool

func main() {

	//variables are defined with var keyword. they are both package and local.
	var i int
	fmt.Println("variable's value:", i, k, l, m)

	//var decleration allows assignment without any type. in below code type of var1 and var2 is integer.
	var var1, var2 = 1, 2
	fmt.Println("var1 and var2:", var1, var2)

	//:= operator can be used instead of var keyword. it is used for decleration and assignment. !!!it is used for only local variables. there has to be var keyword at the beginning of package variables.
	var3 := 7
	fmt.Println("var3:", var3)

	/*
		there are many type in go.
		bool
		string
		int, int8, int16, int32(also known as rune), int64
		uint, uint8(also known as byte), uint16, uint32, uint64, uintptr  **(u means unsigned)
		float32, float64
		complex64, complex128

		!**!int, uint and uintptr types are usually 32 bits wide on 32-bit system, 64 bits wide on 64-bit system!**!

		in format string, %T means type and %v is value for all data type

		each type of variable has its own default value. 0 for numbers, "" for strings, false for bool
	*/

	//multiple variables can be declerad with this way
	var (
		bool_var       bool       = false
		big_number     int64      = 1 << 32
		complex_number complex128 = cmplx.Sqrt(-5 + 12i)
	)

	fmt.Printf("Type: %T, value: %v\n", bool_var, bool_var)
	fmt.Printf("Type: %T, value: %v\n", big_number, big_number)
	fmt.Printf("Type: %T, value: %v\n", complex_number, complex_number)

	//types can be converted to each other with T(v). T->type, v-> value. unlike in C, an integer value is not converted implicitly a float value. in go this operation has to be done explicitly

	var int_value int = 42                       // or int_value := 42
	var float_value float64 = float64(int_value) //or float_value := float64(int_value)
	var unsigned_value uint = uint(float_value)  // or unsigned_value := uint(float_value)
	fmt.Println("unsigned value:", unsigned_value)

	//constants are defined const keyword
	const Pi = 3.14
	fmt.Println("Const:", Pi)

	//numeric constants can take high-precision values || untyped constanst takes the type needed by its context
	const (
		Big   = 1 << 100  //10000000..000000000 ->1 with 100 zeros
		Small = Big >> 99 //2
	)
	fmt.Println("needInt function with Small value:", needInt(Small))
	fmt.Println("needFloat function with Small value:", needFloat(Small))
	fmt.Println("needFloat function with Big value:", needFloat(Big))

}

func needInt(x int) int {
	return 10*x + 1
}

func needFloat(x float64) float64 {
	return 10*x + 1
}
