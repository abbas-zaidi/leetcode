package main

import "fmt"

func main_9() {
	arr := []int{1, 6, 3, 1, 3, 6, 6}
	dups := make([]int, 0)

	n := len(arr)
	for i := 0; i < n; i++ {
		index := arr[i] % n
		arr[index] += n
	}
	fmt.Println(arr)
	for j := 0; j < n; j++ {
		if arr[j]/n > 1 {
			dups = append(dups, j)
		}
	}
	fmt.Println(dups)
}
