package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func BinarySearch[T constraints.Ordered](arr []T, target T) int {
	lpt := 0
	rpt := len(arr) - 1

	for lpt <= rpt {
		mid := (rpt + lpt) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] > target {
			rpt = mid - 1
		} else {
			lpt = mid + 1
		}
	}

	return -1
}

func FirstAndLastIndex[T constraints.Ordered](arr []T, target T) [2]int {
	n := len(arr)
	notFound := [2]int{-1, -1}
	firstPosition := BinarySearch(arr, target)

	if firstPosition == -1 {
		return notFound
	}

	var (
		rightPosition = firstPosition
		leftPosition  = firstPosition
		rightTemp     = rightPosition
		leftTemp      = leftPosition
	)

	for leftPosition > -1 {
		leftTemp = leftPosition
		leftPosition = BinarySearch(arr[0:leftPosition], target)
	}

	for rightPosition > -1 {
		rightTemp = rightPosition
		rightPosition = BinarySearch(arr[(rightPosition+1):n], target)
		if rightPosition > -1 {
			rightPosition = rightTemp + rightPosition + 1
		}
	}

	return [2]int{leftTemp, rightTemp}
}

func binarySearchExample() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	found := BinarySearch(arr, 0)

	fmt.Println(found)
}

func firstAndLastIndexExample() {
	arr := []int{1, 2, 3, 4, 5, 5, 5, 5, 5, 5, 5, 5, 5, 6, 7, 8, 9}

	firstAndLast := FirstAndLastIndex(arr, 5)

	fmt.Println(firstAndLast)
}
