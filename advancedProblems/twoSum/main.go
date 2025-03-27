package main

import "fmt"

func twoSum(nums []int, target int) []int {
	indexMap := make(map[int]int)

	for index, num := range nums {
		complement := target - num
		if complementIndex, found := indexMap[complement]; found {
			return []int{index, complementIndex}
		}
		indexMap[num] = index
	}
	return []int{}
}
func main() {
	nums := []int{4, 7, 5, 1, 2, 3}
	target := 10

	result := twoSum(nums, target)
	fmt.Println(result)
}
