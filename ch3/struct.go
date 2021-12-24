package main

import "fmt"

func main() {
	//define a struct
	type person struct {
		age  int
		name string
	}

	//example of using a struct
	ta := person{
		age:  22,
		name: "tuananh",
	}
	fmt.Println(ta.name)
}
