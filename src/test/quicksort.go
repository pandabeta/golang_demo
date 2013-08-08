// vars project main.go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func partition(array []int, left int, right int) int {

	pivotIndex := left
	pivot := array[left]

	for i := left + 1; i < right+1; i++ {

		if pivot > array[i] {
			pivotIndex += 1

			if pivotIndex != i {
				array[pivotIndex], array[i] = array[i], array[pivotIndex]
			}
		}
	}

	array[left], array[pivotIndex] = array[pivotIndex], array[left]

	return pivotIndex
}

func quicksort(array []int) {

	stack := []int{0, len(array) - 1}

	for len(stack) != 0 {

		size := len(stack)

		left, right := stack[size-2], stack[size-1]

		stack = stack[:size-2]

		pivotIndex := partition(array, left, right)

		if pivotIndex > left {
			stack = append(stack, left, pivotIndex-1)
		}

		if pivotIndex < right {
			stack = append(stack, pivotIndex+1, right)
		}
	}
}

func isSorted(array []int) bool {

	for i := 0; i < len(array)-2; i++ {

		if array[i] > array[i+1] {
			return false
		}
	}

	return true
}

func main() {

	length := 10000000

	list := make([]int, length)

	rand.Seed(time.Now().UTC().UnixNano())

	for i := 0; i < length; i++ {
		list[i] = rand.Intn(length)
	}

	fmt.Println(isSorted(list))

	quicksort(list)

	fmt.Println(isSorted(list))

}

