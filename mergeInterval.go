package main

import (
	"fmt"
)

func main_4() {
	arr := []int{5, 3, 4, 9, 2}
	//arr := [][]int{{2, 4}, {1, 3}, {6, 8}, {1, 10}}
	insertionSort3(arr)
	fmt.Println(arr)
	//arr[0], arr[1] = arr[1], arr[0]
	//mergeIntervals(arr)

}

func mergeIntervals(arr [][]int) {

	n := len(arr)

	//m: = len(arr[0])
	index := 0

	for i := 1; i < n; i++ {

		if arr[index][1] > arr[i][0] {
			arr[index][1] = maxInt(arr[i][1], arr[index][1])

		} else {
			index++
			arr[index] = arr[i]
		}

	}

	fmt.Println("Merged intervals: ")
	for k := 0; k <= index; k++ {
		fmt.Println(arr[k])
	}
}
func maxInt(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func insertionSort(arr [][]int) {
	n := len(arr)

	for i := 1; i < n; i++ {
		j := i

		for j > 0 {
			if arr[j-1][0] > arr[j][0] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
			j = j - 1
		}
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func quickSort(arr []int, low, high int) {
	if low < high {
		p := partition(arr, low, high-1)
		quickSort(arr, low, p-1)
		quickSort(arr, p+1, high-1)
	}
}

func insertionSort2(arr []int) {
	n := len(arr)

	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
			j--
		}
	}
}

func insertionSort3(arr []int) {
	n := len(arr)

	for i := 1; i < n; i++ {
		j := i

		for j > 0 {
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
			j--
		}
	}
}
