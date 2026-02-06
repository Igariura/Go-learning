package main

import "fmt"

func selectionSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {

		smallestIndex := i

		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[smallestIndex] {
				smallestIndex = j
			}
		}
		arr[i], arr[smallestIndex] = arr[smallestIndex], arr[i]
	}
}

func main() {

	numbers := []int{5, 2, 8, 1, 9}
	fmt.Println("Before:", numbers)
	selectionSort(numbers)
	fmt.Println("After:", numbers)
}
