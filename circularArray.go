// You can edit this code!
// Click here and start typing.
package main

import "fmt"

//import "fmt"

func main() {
	nums := []int{-1, -2, -3, -4, -5, 6}

	fmt.Println("circularArrayLoop == ", circularArrayLoop(nums))
}

func circularArrayLoop(nums []int) bool {

	n := len(nums)

	for i := 0; i < n; i++ {
		slow := i
		fast := i
		count := 0

		for {

			slow = nextIndex(slow, nums[slow], n)

			k := nextIndex(fast, nums[fast], n)
			fast = nextIndex(k, nums[k], n)

			if fast == k {
				break
			}
			fmt.Printf("slow =%v, fast = %v, count = %v\n", slow, fast, count)

			if slow == fast && count > 1 {
				return true
			}
			if slow != fast {
				count++
			}

		}

	}

	return false
}

func nextIndex(j int, k int, n int) int {
	z := (j + k) % n

	if z < 0 {
		z = z + n
	}

	return z
}
