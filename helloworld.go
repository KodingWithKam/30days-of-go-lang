/*
A Go file consists of:

Package declaration
Imported packages
Functions
Statements and expressions
*/

package main

import (
	"fmt"
)

func main() {
	/*
		GO SLICES:
		Slices are similar to arrays, but are more powerful and offer more flexibility.
		Like arrays, slices are also used to store multiple values of the same type in a single variable.
		However, unlike arrays, the length of a slice can grow and shrink as you see fit.

		In GO there are a few ways to make a slice:
		1. Using the []datatype{values} format 👈🏾 slice_name := []datatype{values}

		2. Create a slice from an array
			var myArray = [length]datatype{values} // An array
			mySlice := myArray[start:end] // A slice made from the array

		3. Using the make() function 👈🏾 mySlice1 := make([]int, 5, 10)
	*/

	// Example 1 slice with []datatype{values}
	slice1 := []int{1, 2, 3} // 👈🏾declares a slice of integers of length 3 and also the capacity of 3 ✅
	fmt.Println(slice1)

	/*
		LENGTH & CAPACITY
		In Go, there are two functions that can be used to return the length and capacity of a slice:

		len() function - returns the length of the slice (the number of elements in the slice)
		cap() function - returns the capacity of the slice (the number of elements the slice can grow or shrink to)
	*/

	// Example 1 getting length and capacity values from slice
	slice2 := []string{"Go", "Slices", "Are", "Powerful"}
	fmt.Println(len(slice2)) // 👈🏾will print 4 ✅
	fmt.Println(cap(slice2)) // 👈🏾will print 4 ✅
	fmt.Println(slice2)      // 👈🏾will print [Go Slices Are Powerful] ✅

	// Example 2 create slice from array
	arr1 := [6]int{10, 11, 12, 13, 14, 15}
	slice3 := arr1[2:4] // a slice with length 2. It is made from arr1 which is an array with length 6.
	// The slice starts from the second element of the array which has value 12. The slice can grow to the end of the array.
	// This means that the capacity of the slice is 4.
	fmt.Printf("slice3 = %v\n", slice3)        // 👈🏾will print slice3 = [12,13] ✅
	fmt.Printf("length = %d\n", len(slice3))   // 👈🏾will print length = 2 ✅
	fmt.Printf("capacity = %d\n", cap(slice3)) // 👈🏾will print capacity = 4 ✅

	// Example 3 create slice using make() function
	slice4 := make([]int, 5, 10)               //👈🏾 If the capacity parameter is not defined, it will be equal to length.
	fmt.Printf("slice4 = %v\n", slice4)        // 👈🏾will print slice4 = [0 0 0 0 0] ✅
	fmt.Printf("length = %d\n", len(slice4))   // 👈🏾will print length = 5 ✅
	fmt.Printf("capacity = %d\n", cap(slice4)) // 👈🏾will print capacity = 10 ✅

	// Example 4 create slice using make() function and no capacity specified
	slice5 := make([]int, 5)
	fmt.Printf("slice5 = %v\n", slice5)        // 👈🏾will print slice5 = [0 0 0 0 0] ✅
	fmt.Printf("length = %d\n", len(slice5))   // 👈🏾will print length = 5 ✅
	fmt.Printf("capacity = %d\n", cap(slice5)) // 👈🏾will print capacity = 5 ✅
}
