package main

import (
	"fmt"
)

func main() {
	// first declaration of an array
	var myArray [3]int

	// declaration of an array with make
	myArray2 := make([]int, 4)

	//declaration of an array with make adding limit
	// myArray3 := make([]int, 4, 5)
	myArray3 := [4]int{1, 2, 3, 4}

	fmt.Println("the length of array1: ", len(myArray))
	fmt.Println("the length of array2: ", len(myArray2))
	fmt.Println("the length of array3: ", len(myArray3))

	fmt.Println("the capacity of array1: ", cap(myArray))
	fmt.Println("the capacity of array2: ", cap(myArray2))
	fmt.Println("the capacity of array3: ", cap(myArray3))
	// for i:= range myArray3 {
	// 	fmt.Print("Enter value at index ", i, ": ")
	// 	var temp int
	// 	_, err := fmt.Scan(&temp)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	myArray3[i] = temp
	// }

	fmt.Println(myArray3)

	mySlice := myArray3[:1]
	// mySlice2 := make([]int, 2)
	fmt.Println(mySlice)
	mySlice[0] = 33
	fmt.Println(mySlice)
	fmt.Println(myArray3)

}
