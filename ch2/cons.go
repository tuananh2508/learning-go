package main

import "fmt"

// Declared constant outside of function
const y int = 1

func main() {
	// Error will appeared when you try to do this
	//y = y + 1
	//fmt.Println(y)

	// Constant can also be declared inside of a function
	const x = "Hello"
	fmt.Println(x)
	// Again if you change the value of x It will show an Error
	//x = "bye"
	//fmt.Println(x)
}
