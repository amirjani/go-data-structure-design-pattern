// Brute Force algorithms are known for wide applicability and simplicity in solving complex problems.
// Searching, string matching, and matrix multiplication are some scenarios where they are used
//main package has examples shown
//in Hands-On Data Structures and algorithms with Go book
package main

// importing fmt package
import (
	"fmt"
)

//findElement method given array and k element
func findElement(arr [10]int, k int) bool {
	var i int
	for i = 0; i < 10; i++ {
		if arr[i] == k {
			return true
		}
	}
	return false
}

// main method
func main() {
	var arr = [10]int{1, 4, 7, 8, 3, 9, 2, 4, 1, 8}
	var check bool = findElement(arr, 10)
	fmt.Println(check)
	var check2 bool = findElement(arr, 9)
	fmt.Println(check2)
}
