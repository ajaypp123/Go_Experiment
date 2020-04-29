/*
Slices:
Similar to Array but like dynamic array were capacity can b change.
Array have length but slices have length and capacity.
In Slices initital size set to size of array and after append it increment by 2*size
*/
package main

import "fmt"

func mySlices() {
	/*
		Assign
		Length: 3 Capacity: 100, Capacity parameter is optional
		s := make([]int, 3, 100)
	*/
	//s := make([]int, 3)
	//s[0] = 22
	//s[1] = 33
	//s[2] = 44
	s := []int{22, 33, 44}

	/*
		append
	*/
	s = append(s, 55, 66)
	fmt.Println(s[3:], s[:3], s[:], len(s), cap(s))

	/*
		copy
	*/
	// Here both will get same pointer so, change in s will change in x
	x := s                   // shallow copy
	y := make([]int, len(x)) // deep copy
	copy(y, s)
	s[0] = 500
	fmt.Println(x[3:], x[:3], x[:])
	fmt.Println(s[3:], s[:3], s[:])
	fmt.Println(y[3:], y[:3], y[:])

	/*
		Myltidimentional with dynamic array inside
		[[2], [33,44], [55,66,77]]
	*/
	mulArr := make([][]int, 3)
	for i := 0; i < len(mulArr); i++ {
		mulArr[i] = make([]int, i+1)
		for j := 0; j < len(mulArr[i]); j++ {
			mulArr[i][j] = j + 1
		}
	}
	fmt.Println(mulArr)
}

/*
func main() {
	mySlices()
}
*/
