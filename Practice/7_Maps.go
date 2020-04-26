// Maps are similar to list of python
package main

import "fmt"

func myMaps() {
	/*
		Declare
	*/
	//dict := make(map[string]int)
	//dict["a"] = 1
	//dict["b"] = 2
	dict := map[string]int{"a": 1, "b": 2}
	fmt.Println(dict)

	/*
		Delete Element
	*/
	delete(dict, "a")
	fmt.Println(dict)

	/*
		Check Element in map
	*/
	val, isPresent := dict["a"]
	fmt.Println(val, isPresent)

	val1, isPresent1 := dict["b"]
	fmt.Println(val1, isPresent1)
	fmt.Println(dict["c"])
}
