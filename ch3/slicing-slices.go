package main

import "fmt"

func main() {
	//slices
	x := []int{1, 2, 3, 4}
	fmt.Println("original slice:", x)
	//slicing the slices
	y := x[:2]
	z := x[1:]
	k := x[1:3]
	m := x[:]
	fmt.Println("after slicing: ", y, z, k, m)

	//slices share memory - so changes in sub-slices gonna affect the original slice.
	//capacity of subslices is going to the same as the original but minus the subset. For example: cap(x)=5 , with y := x[2:] this mean cap of y gonna be 5-2=3.

	j := []int{1, 2, 3, 4}
	i := j[:2]
	fmt.Println(cap(j), cap(i))
	i = append(i, 30)
	fmt.Println(j, i)

	// NOTE: never append to subslices, if you must do so, use full slice expression: n := m[low:high:max]
	n := i[0:1:3]
	fmt.Println(n)

}
