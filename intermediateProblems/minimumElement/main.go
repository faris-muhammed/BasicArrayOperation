package main

import "fmt"

func minimum(array []int) (int, error) {
	if len(array) == 0 {
		return 0, fmt.Errorf("array is empty")
	}
	minimum := array[0]
	for _, v := range array {
		if minimum > v {
			minimum = v
		}
	}
	return minimum, nil
}
func main() {
	array := []int{10, 4, -7, 51, 6, -1, 1, 2}

	result, err := minimum(array)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Minimum element in the array:", result)
	}
}
