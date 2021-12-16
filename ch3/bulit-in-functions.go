package main

import "fmt"

func main() {
	//len function
	var x = [3]int{1, 2, 3}
	fmt.Println("Lenght of X is: ", len(x))

	//append function for slices
	var y []int
	fmt.Println(y)
	y = append(y, 4, 4, 5, 5)
	fmt.Println(y)

	//append slices to slices
	j := []int{4, 5, 6, 7}
	i := []int{1, 2, 3, 4}
	i = append(i, j...)
	fmt.Println(i)

	// findout capacity of slice
	k := []int{1, 2, 3, 10: 100}
	fmt.Println(k)
	fmt.Println("Capacity of slice: ", cap(k))
	// How capacity worked: https://go.dev/play/p/qh_hyXLbgkM

}
