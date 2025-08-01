package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 2, 4, 5, 3, 6, 7, 1}

	duplicateNumbers := findDuplicate(arr)
	fmt.Println(duplicateNumbers)
}

func findDuplicate(arr []int) []int {
	occurences := make(map[int]int)
	result := []int{}

	for _, value := range arr {
		occurences[value]++

		if occurences[value] == 2 {
			result = append(result, value)
		}
	}

	return result
}
