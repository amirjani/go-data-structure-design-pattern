// O(n)
package main

import "fmt"

func main() {
	var m [10]int
	var k int

	for k = 0; k < 10; k++ {
		m[k] = k * 100
		fmt.Printf("Element[%d] = %d\n", k, m[k])
	}
}
