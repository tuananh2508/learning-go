package main

import "fmt"

func main() {
	// slices are dynamic while array is not
	var x = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(x)

	// 2 dimensional
	var y = [][]int{
		{1, 2, 3, 4, 5, 6, 7, 8},
		{1, 2, 3, 4, 5, 6, 7, 9},
	}
	fmt.Println(y)

	//nill slices
	var k []int
	fmt.Println(k)

}
