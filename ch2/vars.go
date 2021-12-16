package main

import "fmt"

//specify a variable outside of function using var
// := declaration is not supported outside of function
//var x = 10

func main() {
	var x = 20
	y := 20
	//can't use var the change the value of variables
	//var x = 20

	//way to define multiple variables
	//var x,y = 10, 20
	// var (
	// 	x=10;
	//  y=1;
	//)

	//in the other side, := can do so as long as the left side has new variables (at least 1)
	//x, z := 30, 30
	//fmt.Println(x, y, z)

	// Typed and Untyped Constants (Below are untyped constant)
	const k = 30
	var z float64 = k + 1.2
	fmt.Println(x, y, z, k)

	// Unused var. Compipler wont complain about this
	var unused_var = "Hello"
	unused_var = "World"
	fmt.Println(unused_var)
	// Also not complain about this because the variable is used before
	unused_var = "!"
	// But golangci-lint gonna know about this. So take care.
	// Unused constant is not going to be compile.
	const constant_used = "never"

}
