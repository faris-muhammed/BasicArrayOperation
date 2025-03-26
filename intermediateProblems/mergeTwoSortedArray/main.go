package main

import "fmt"

func mergeSortedArray(arr1 []int, arr2 []int) []int {
	i, j := 0, 0
	var result []int
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] <= arr2[j] {
			result = append(result, arr1[i])
			i++
		} else {
			result = append(result, arr2[j])
			j++
		}
	}
	for i < len(arr1) {
		result = append(result, arr1[i])
		i++
	}
	for j < len(arr2) {
		result = append(result, arr2[j])
		j++
	}
	return result
}
func main() {
	array1 := []int{1, 3, 5, 7}
	array2 := []int{2, 4, 6, 8}

	fmt.Println("After merge two sorted array", mergeSortedArray(array1, array2))
}
