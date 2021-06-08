// Lists have a variable length and developer can remove or add elements more easily than an array
// Data items within a list need not be contiguous in memory or on disk
package main

import (
	"container/list"
	"fmt"
)

func main() {
	var intList list.List
	intList.PushBack(1)
	intList.PushBack(2)
	intList.PushBack(3)

	for element := intList.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value.(int))
	}
}
