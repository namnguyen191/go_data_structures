package main

import (
	"fmt"
)

func main() {
	firstAndLastIndexExample()
}

func heapExample() {
	minIntegerHeap := NewMinIntHeap([]int{8, 3, 2, 1, 3, 5, 6, 3, 0, 1, 2, 3, 11, 22, 33})

	fmt.Println(minIntegerHeap)

	for {
		num, err := minIntegerHeap.Poll()
		if err != nil {
			break
		}
		fmt.Println(num)
	}
}
