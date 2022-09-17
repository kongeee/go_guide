package main

import (
	"fmt"
	"strings"
)

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
	names := [5]string{
		"büşra",
		"furkan",
		"eren",
		"yasir",
        "civan",
	}

	fmt.Println("names before change:", names)
	slice1 := names[0:2] //büşra, furkan
	slice2 := names[1:3] //furkan, eren
    slice3 := names[5:]  //civan
    

	slice2[0] = "XXXXXXXX" //before : slice2[0] -> furkan, after : slice2[0] -> XXXXXXXX . this change affects original array and other slices
	fmt.Printf("Original array :%v | slice1 :%v | slice2 :%v\n", names, slice1, slice2)
    fmt.Printf("Critical slice :%v", slice3)

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

	//A slice has both a length and a capacity. length means number of elements that the slice contains. capacity means total size of the slice.
	slice_x := []int{3, 5, 7, 9, 11, 13, 15, 17}
	fmt.Printf("length: %v | capacity: %v\n", len(slice_x), cap(slice_x))
	s1 := slice_x[:4]
	fmt.Printf("length: %v | capacity: %v\n", len(s1), cap(s1)) //capacity is 8. because it can still hold 8 elements
	s2 := slice_x[2:5]
	fmt.Printf("length: %v | capacity: %v\n", len(s2), cap(s2)) //capacity is 6. because two elements at the beginning are dropped

	//there is no null in go. insteak of it, there is nil.
	var slice4 []int
	if slice4 == nil { //capacity and lenght are 0
		fmt.Println("slice4 is nil", slice4, len(slice4), cap(slice4))
	}

	//built-in make() function can be used for creating a slice. if lenght is not 0, 0 is put as the length number in the slice

	k := make([]int, 5) //len(k) = 5 , cap(k) = 5 and k = [0 0 0 0 0]
	fmt.Printf("len = %v | cap = %v | k = %v\n", len(k), cap(k), k)

	m := make([]int, 0, 5) //len(n) = 0 , cap(m) = 5 and k = []
	fmt.Printf("len = %v | cap = %v | k = %v\n", len(m), cap(m), m)
	//!! length can be increased until it reaches capacity value
	t := m[1:4]
	fmt.Println("slice t:", t)

	//slices can contain slice -> n-dimension array

	board := [][]string{
		[]string{"-", "-", "-"},
		[]string{"-", "-", "-"},
		[]string{"-", "-", "-"},
	}
	board[0][0] = "X"
	board[0][1] = "O"
	board[1][1] = "X"
	board[2][1] = "O"
	board[2][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	//n item is added to a slice by using append() function
	var n []int
	z := []int{7, 8, 9}
	fmt.Printf("len = %v | cap = %v\n", len(n), cap(n))

	n = append(n, 0)
	fmt.Printf("len = %v | cap = %v\n", len(n), cap(n))

	n = append(n, 1, 2, 3, 4)
	fmt.Printf("len = %v | cap = %v\n", len(n), cap(n))

	n = append(n, z...)
	fmt.Printf("len = %v | cap = %v\n", len(n), cap(n)) /*??????????????? i cannot understand why length is 8 while capacity is 12 ????????????????????????????????????????????????????*/

}
