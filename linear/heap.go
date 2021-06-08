// The heap data structure is used in selection, graph, and k-way merge algorithm
// Operations such as finding, merging, insertion, key changes, and deleting are performed on heaps
// If the order is descending, it is referred to as a maximum heap; otherwise, it's a minimum heap
// It is not a sorted data structure, but partially ordered.
package main

import (
	"container/heap"
	"fmt"
)

type IntegerHeap []int

func (iHeap IntegerHeap) Len() int {
	return len(iHeap)
}

func (iHeap IntegerHeap) Less(firstElement, secondElement int) bool {
	return iHeap[firstElement] < iHeap[secondElement]
}

func (iHeap IntegerHeap) Swap(firstElement, secondElement int) {
	iHeap[firstElement], iHeap[secondElement] = iHeap[secondElement], iHeap[firstElement]
}

func (iHeap *IntegerHeap) Push(heapIntFormat interface{}) {
	*iHeap = append(*iHeap, heapIntFormat.(int))
}

func (iHeap *IntegerHeap) Pop() interface{} {
	var heapLength int
	var lastNode int
	var previous IntegerHeap = *iHeap

	heapLength = len(previous)
	lastNode = previous[heapLength-1]
	*iHeap = previous[0 : heapLength-1]

	return lastNode
}

func main() {
	var intHeap *IntegerHeap = &IntegerHeap{1, 3, 4}

	heap.Init(intHeap)
	heap.Push(intHeap, 2)

	fmt.Printf("minimum: %d\n", (*intHeap)[0])
	for intHeap.Len() > 0 {
		fmt.Printf("%d \n", heap.Pop(intHeap))
	}
}
