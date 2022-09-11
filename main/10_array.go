package main

import "fmt"

func main() {

	//array defination -> [n]T , n->number of value, T->type
	//array size is fixed
	var a [10]int
	a[0] = 1
	a[2] = 65
	fmt.Println("a array:", a)

	//array can be sliced. array[low : high]. slices are much more common than arrays!!!!!!!!!!!
	primes := []int{2, 3, 5, 7, 11}
	var s []int = primes[1:4]
	fmt.Println("slice:", s)

	//slices are like references to arrays. when a slice that is generated from an array is changed, this change affects original array and other slices that is generated from the same array
	names := [4]string{
		"büşra",
		"furkan",
		"eren",
		"yasir",
	}

	fmt.Println("names before change:", names)
	slice1 := names[0:2] //büşra, furkan
	slice2 := names[1:3] //furkan, eren

	slice2[0] = "XXXXXXXX" //before : slice2[0] -> furkan, after : slice2[0] -> XXXXXXXX . this change affects original array and other slices
	fmt.Printf("Original array :%v | slice1 :%v | slice2 :%v\n", names, slice1, slice2)

	//array and slice definations are a little bit different. Ex. [3]bool{true,true,false} -> array |||| []bool{true,true,false} -> slice
	//a struct array or slice can be defined like this

	s_arr := []struct {
		i int
		b bool
	}{
		{2, true},
		{4, false},
		{7, false},
	}
	fmt.Println("struct slice:", s_arr)

	//when slicing low part or high part may be omitted
	numbers := []int{2, 3, 5, 7, 11, 13, 17}
	num_slice1 := numbers[:4]
	num_slice2 := numbers[2:]
	num_slice3 := numbers[:]

	fmt.Printf("slice default : normal = %v | [:4] = %v | [2:] = %v | [:] = %v\n", numbers, num_slice1, num_slice2, num_slice3)

	//

}
