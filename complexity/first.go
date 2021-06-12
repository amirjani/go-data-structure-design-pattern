// An algorithm is a set of steps to be processed by different operations to achieve a task
package main

import "fmt"

func main() {
	var m [10]int
	var k int
	sum := 0

	for k = 0; k < 10; k++ {
		m[k] = k + 200
		sum += m[k]
		fmt.Printf("Element[%d] = %d\n", k, m[k])
	}
	fmt.Println("sum is: ", sum)
}
