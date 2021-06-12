package main

import "fmt"

func main() {
	var arr [10][10][10]int

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			for k := 0; k < 10; k++ {
				arr[i][j][k] = i + j + k
			}
		}
	}

	fmt.Println(arr)
}
