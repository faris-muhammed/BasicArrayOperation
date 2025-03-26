package main

import "fmt"

func rotate(array []int, k int) []int {
	n := len(array)
	if n == 0 {
		return array
	}
	k = k % n
	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[(i+k)%n] = array[i]
	}
	return result
}

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	result := rotate(array, k)
	fmt.Println("Right rotated array:", result)
}
