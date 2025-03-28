package main

import "fmt"

func missingNumber(array []int, n int) int {
	expectedSum := n * (n + 1) / 2
	actualSum := 0

	for _,v := range array {
		actualSum += v
	}
	return expectedSum - actualSum
}

func main() {
	nums := []int{1, 2, 3, 4, 6}
	n := 6
	fmt.Println("The missing number :", missingNumber(nums, n))
}
