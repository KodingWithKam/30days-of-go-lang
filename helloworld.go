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
		Arrays are used to store multiple values of the same type in a single variable,
		instead of declaring separate variables for each value.

		There are two ways to declare an array:

		1. *USING THE VAR KEYWORD*
		var array_name = [length]datatype{values} <-- here length is defined
			or
		var array_name = [...]datatype{values} <-- here length is inferred

		2. *USING THE := SYMBOL*
		array_name := [length]datatype{values} <-- here length is defined
			or
		array_name := [...]datatype{values} <-- here length is inferred

		** The length specifies the number of elements to store in the array.
		In Go, arrays have a fixed length. The length of the array is either defined by a number or
		is inferred (means that the compiler decides the length of the array, based on the number of values).
	*/

	// Example 1 using var keyword with defined lengths ✅
	var arr1 = [3]int{1, 2, 3}
	var arr2 = [4]string{"one", "two", "three", "four"}

	// Example 2 using var keyword with inferred lengths ✅
	var arr3 = [...]int{1, 2}
	var arr4 = [...]string{"one", "two", "three"}

	// Example 1 using := symbol with defined lengths ✅
	arr5 := [3]int{1, 2, 3}
	arr6 := [2]string{"one", "two"}

	// Example 2 using := symbol with inferred lengths ✅
	arr7 := [...]int{1, 2}
	arr8 := [...]string{"one", "two", "three"}

	/*
		ACCESSING ARRAY ELEMENTS 👁
		You can access a specific array element by referring to the index number.
		In Go, array indexes start at 0. That means that [0] is the first element, [1] is the second element, etc.
	*/

	// Example 1 Accessing array elements
	var arr9 = [...]int{10, 100, 1000}
	fmt.Println(arr9[1]) // <-- Will print the second element 💯since array indexes start from 0 ✅
	fmt.Println(arr9[0]) // <-- Will print the first element 10 since array indexes start from 0 ✅

	/*
		CHANGING ARRAY ELEMENT VALUES
		You can also change the value of a specific array element by referring to the index number.
	*/

	// Example 1 Changing array element values
	var arr10 = [...]int{1, 2, 3}
	arr10[2] = 40
	fmt.Println(arr10) // <-- Will print out [1 2 40] ✅

	/*
			ARRAY INITIALIZATION 🫀
			If an array or one of its elements has not been initialized in the code,
			it is assigned the default value of its type.

		 	The default value for int is 0, and the default value for string is "".
	*/

	// Example 1 ARRAY INIT
	arr11 := [5]int{}              //not initialized
	arr12 := [5]int{1, 2}          //partially initialized
	arr13 := [5]int{1, 2, 3, 4, 5} //fully initialized

	fmt.Println(arr11) // 👈🏾will print [0 0 0 0 0] ✅
	fmt.Println(arr12) // 👈🏾 will print [1 2 0 0 0] ✅
	fmt.Println(arr13) // 👈🏾 will print [1 2 3 4 5] ✅

	// Example 2 ARRAY INIT SPECIFIC VALUES
	arr14 := [5]int{1: 10, 2: 40}

	fmt.Println(arr14) // 👈🏾 will print [0 10 40 0 0] ✅

	/*
		LENGTH OF AN ARRAY
		The len() function is used to find the length of an array
	*/

	// Example 1 Array length
	arr15 := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(len(arr15)) // 👈🏾 will print 6 ✅

	// Simply prints out contents
	fmt.Println(arr1, arr2, arr3, arr4, arr5, arr6, arr7, arr8)
}
