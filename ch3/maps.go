package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["n1"] = 1
	m["n2"] = 2
	a := m["n1"]
	b := m["n2"]
	fmt.Println(a, b)

	teams := map[string][]string{
		"Orcas":   []string{"Fred", "Ralph", "Bijou"},
		"Lions":   []string{"Sarah", "Peter", "Billie"},
		"Kittens": []string{"Waldo", "Raul", "Ze"},
	}
	team1 := teams["Orcas"]
	team2 := teams["Lions"]
	team3 := teams["Kittens"]

	// ",ok" used to check if the variable is present in the map or not
	team4, ok := teams["Holland"]
	fmt.Println(team1, team2, team3, team4, ok)

	// Delete element from map
	number := map[string]int{
		"one": 1,
		"two": 2,
	}
	delete(number, "one")
	fmt.Println(number)

}
