package main

import "fmt"

func Insert(arr []int, index, element int) ([]int, error) {
	if index < 0 || index > len(arr) {
		return arr, fmt.Errorf("invalid index")
	}
	arr = append(arr[:index+1], arr[index:]...)
	arr[index] = element
	return arr, nil
}

func Delete(arr []int, index int) ([]int, error) {
	if index < 0 || index >= len(arr) {
		return arr, fmt.Errorf("invalid index")
	}
	arr = append(arr[:index], arr[index+1:]...)
	return arr, nil
}

func Search(arr []int, element int) (int, error) {
	for k, v := range arr {
		if v == element {
			return k, nil
		}
	}
	return -1, fmt.Errorf("no such value")
}

func Reverse(arr []int) ([]int, error) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr, nil
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}

	// Insert
	index := 2
	value := 100
	arr, err := Insert(arr, index, value)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("After insert to array", arr)
	}

	// Delete
	arr, err = Delete(arr, index)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("After deletion", arr)
	}

	// Search
	element := 1
	searchIndex, err := Search(arr, element)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("The element is at position", searchIndex)
	}

	// Reverse
	fmt.Println("Before reverse", arr)
	arr, err = Reverse(arr)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("After reversed", arr)
	}
}
