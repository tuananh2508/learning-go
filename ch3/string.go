package main

import "fmt"

func main() {
	// go use byte to represent string
	var s string = "Hello World"
	var b byte = s[6]
	fmt.Println(s, b)

	// string can do same things as slice
	var s2 string = s[:5]
	var s3 string = s[3:]
	var s4 string = s[3:4]
	fmt.Println(s2, s3, s4, len(s2))

	// but string with other languages can cause some problems

	//integer or byte can convert to string
	i := int(65)
	si := string(i)
	fmt.Println(si)
	var by = byte(67)
	sby := string(by)
	fmt.Println(sby)
}
