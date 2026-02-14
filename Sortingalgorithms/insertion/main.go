package main

import "fmt"

func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}

		arr[j+1] = key
	}

}

func main() {

	arr := []int{8, 3, 7, 1, 5}

	fmt.Println("Original array:", arr)

	insertionSort(arr)

	fmt.Println("Sorted array: ", arr)
}
