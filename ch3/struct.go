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
	fmt.Println(ta.name, ta.age)

	// can compare the same type of data with different struct
	// defind first struct
	type first_person struct {
		name string
		age  int
		job  string
	}
	//sample values
	bob := first_person{
		name: "bob",
		age:  22,
		job:  "driver",
	}
	//define sencond struct
	type second_person struct {
		name string
		age  int
		job  string
	}
	//sample values
	alice := second_person{
		name: "alice",
		age:  23,
		job:  "teacher",
	}
	//compare same type of data with 2 structs
	fmt.Println(bob.age != alice.age)
}
