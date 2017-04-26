package main

import "fmt"

func collections() {
	//Array
	myArray := [3]int{}
	myArray[0] = 5
	myArray[1] = 10
	myArray[2] = 20
	fmt.Println(myArray)

	//Array
	myArray2 := [...]int{5, 10, 20}
	fmt.Println(myArray2)

	//Slice
	mySlice := myArray2[:]
	fmt.Println(mySlice)
	fmt.Println(len(mySlice))

	//Map
	myMap := make(map[int]string)
	fmt.Println(myMap)

	myMap[3] = "Foo"
	myMap[9] = "Bar"
	fmt.Println(myMap)
}

// Prints
// [5, 10, 20]
// [5, 10, 20]
// [5, 10, 20]
// 3
// map[]
// map[9:Bar 3:Foo]
