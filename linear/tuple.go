// A tuple is a finite sorted *list* of elements
// It is a data structure that groups data
package main

import "fmt"

// * this return a tuple
func powerSeries(a int) (int, int) {
	return a * a, a * a * a
}

func main() {
	var square int
	var cube int

	square, cube = powerSeries(3)

	fmt.Println("Square: ", square, "Cube: ", cube)
}
