package main

import "fmt"

func main() {
	//copy function take 2 input, first one is the destination slice, 2nd is the source slice/string
	x := []int{1, 2, 3, 4}
	y := make([]int, 2)
	copy(y, x)
	z := make([]int, 2)
	copy(z, x[2:])
	fmt.Println(x, y, z)

}
