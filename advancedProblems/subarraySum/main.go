package main

import "fmt"

func subarraySum(nums []int, k int) int {
	prefixSum := make(map[int]int)
	prefixSum[0] = 1

	sum := 0
	count := 0

	for _, num := range nums {
		sum += num
		if val, ok := prefixSum[sum-k]; ok {
			count += val
		}
		prefixSum[sum] += 1
	}
	return count
}

func main() {
	array := []int{2, 4, 7, 3, 3, 6}
	k := 6
	fmt.Println("Total subarrays with sum", k, ":", subarraySum(array, k))
}
