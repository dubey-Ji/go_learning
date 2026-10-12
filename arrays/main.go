package main

import (
	"fmt"
	"sort"
)

func main() {
	// ** Declare array
	// ** Syntax 1
	// var arr [4]int
	// arr[0] = 0
	// arr[1] = 1
	// arr[2] = 2
	// arr[3] = 3
	// fmt.Println(arr)

	// ** Syntax 2
	arr := [4]int{1, 3, 2, 4}
	// fmt.Println(arr)

	// ** Loop through an array
	// ** for loop
	// fmt.Println("simple for loop")
	// for i := 0; i < len(arr); i++ {
	// 	fmt.Println(arr[i])
	// }

	// ** for loop with range
	// fmt.Println("For loop with range")
	// for _, val := range arr {
	// 	fmt.Println(val)
	// }

	// ** Copy one array to another array
	/*
		for this we have a copy() function which takes argument as destination and source
		and return the number of element copied
	*/
	// var copyArr [5]int
	// fmt.Println("copyarr before", copyArr)
	// n := copy(copyArr[:], arr[:])
	// fmt.Println("number of element copied", n)
	// fmt.Println("copy array", copyArr)

	// ** Sorting array
	// ** Ascending order sort
	sort.Ints(arr[:]) // it sort the array that is paced in place no return
	// fmt.Println(arr)

	// ** Array comparison
	arr2 := [4]int{1, 2, 4, 3}
	arr3 := [4]int{1, 2, 3, 4}
	fmt.Println(arr2 == arr3)
	fmt.Println(arr == arr3)

	// ** Check if element is present or not.
	// ** Can use loop here
	var searchNumber int = 3
	for idx, val := range arr2 {
		if val == searchNumber {
			fmt.Printf("Number %v found at index %v \n", searchNumber, idx)
		}
	}

}
