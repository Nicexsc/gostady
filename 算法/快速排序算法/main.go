package main

import "fmt"

func QuickSort(arr []int) []int {
	count := len(arr)
	for i := 0; i < count; i++ {
		for j := i + 1; j < count; j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}

		}

	}
	return arr

}

func main() {
	arr := []int{1, 9, 10, 30, 2, 5, 45, 8, 63, 234, 12}
	fmt.Println(QuickSort(arr))
}