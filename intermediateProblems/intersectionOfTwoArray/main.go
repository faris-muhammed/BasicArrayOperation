package main

import "fmt"

func findIntersection(arr1 []int, arr2 []int) []int {
	myMap := make(map[int]bool)
	result := []int{}

	for _, v := range arr1 {
		myMap[v] = true
	}

	for _, v := range arr2 {
		if myMap[v] {
			result = append(result, v)
			delete(myMap, v)
		}
	}
	return result
}

func main() {
	arr1 := []int{1, 2, 2, 1, 3, 4}
	arr2 := []int{2, 2, 3, 5}
	fmt.Println("Intersection:", findIntersection(arr1, arr2))
}
