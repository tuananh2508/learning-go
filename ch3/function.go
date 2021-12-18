package main

import "fmt"

func main() {
	// append function: append elements at the end of slices (only work with slices)
	var x = []int{1, 2, 3, 4}
	x = append(x, 1, 2)
	fmt.Println(x)

	y := []int{1, 2, 3, 4}
	y = append(y, 1, 2)
	fmt.Println(y)

	//cap function: show capacity of slices/array
	fmt.Println(cap(x))

	//make function: make a "blank" slices. 1st input - type , 2nd input - len , 3rd input - cap
	z := make([]int, 5, 10)
	z = append(z, 1)
	fmt.Println(z, len(z), cap(z))

	//empty vs nil slices
	var j = make([]int, 1)
	var i []int
	fmt.Println(i == nil, j == nil)

	//make create a empty slices - not nil slice
	var k = make([]int, 1, 1)
	fmt.Println(len(k), k == nil)

}
