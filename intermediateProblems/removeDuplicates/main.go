package main

import (
	"fmt"
)

func removeDuplicate(array []int) []int {
	myMap := make(map[int]bool)
	var result []int

	for _, v := range array {
		if !myMap[v] {
			myMap[v] = true
			result = append(result, v)
		}
	}
	return result
}
func main() {
	array := []int{1, 2, 2, 3, 3, 45, 45}
	res := removeDuplicate(array)
	fmt.Println(res)
}
