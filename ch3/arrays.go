package main

import "fmt"

func main() {
	// Normal Array
	var x = [3]int{10, 20, 30}
	fmt.Printf("X variables are: %d %d %d\n", x[0], x[1], x[2])
	// Sparse Array
	y := [12]int{1, 5: 4, 6, 10: 100, 14}
	fmt.Println("Values of elements in array Y are:", y)

	// Can specify and array without number of elements in array
	z := [...]float64{0, 1, 2, 3}
	fmt.Println("Values of elements in array Z are (without specify number of elements):", z)

	// 2 dimensional
	var k = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("Array K are 2 dimensional:\n", k[0][0], k[0][1], k[0][2], "\n", k[1][0], k[1][1], k[1][2])

	// Cant read or write past the end of the array
	//j := [1]int{0}
	//fmt.Println(j[1])

	//Lenght of an array
	fmt.Println("Lenght of array x is: ", len(x))

	// using constant variable to define lenght of arrays. But the normal variable not gonna work
	const j = 2
	var i = [j]int{1, 2}
	fmt.Println(i)
}
