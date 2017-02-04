package main

import "fmt"

func main () {
	var sessions = map[string]int{
		"foo": 1,
		"bar": 2,
	};
	fmt.Println(sessions)
	delete(sessions, "foo");
	fmt.Println(sessions)



	pays := map[string][]string{
		"pepe": {"1", "2", "3"},
	}

	fmt.Println(pays["pepe"])

	fmt.Println(pays)
}
