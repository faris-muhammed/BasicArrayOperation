package main

import "fmt"

func maximum(array []int) (int, error) {
	if len(array) == 0 {
		return 0, fmt.Errorf("array is empty")
	}
	maximum := array[0]
	for _, v := range array {
		if maximum < v {
			maximum = v
		}
	}
	return maximum, nil
}
func main() {
	array := []int{-10, -4, -7, -50, -2, -12}

	result, err := maximum(array)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Maximum element in this array:", result)
	}
}
