package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{2, 3, 7, 6, 8, -1, -10, -15}

	min := math.MaxInt

	for i := 0; i < len(arr); i++ {
		if arr[i] > 0 && arr[i] < min {
			min = arr[i]
		}
	}
	fmt.Println("Min = ", min)
}
