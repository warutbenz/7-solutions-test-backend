package main

import "fmt"

func main() {

	input := [][]int{
		{59},
		{73, 41},
		{52, 40, 53},
		{26, 53, 6, 1000},
	}
	maximumpath := findMaximumPath(input)

	fmt.Println(input)
	fmt.Println(maximumpath)
}

func findMaximumPath(input [][]int) int {

	n := len(input)
	// copy data to temp
	arr := make([][]int, n)
	for i := range arr {
		arr[i] = make([]int, len(input[i]))
		copy(arr[i], input[i])
	}
	// find max path
	for i := len(arr) - 2; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			if arr[i+1][j] > arr[i+1][j+1] {
				arr[i][j] += arr[i+1][j]
			} else {
				arr[i][j] += arr[i+1][j+1]
			}
		}
	}
	return arr[0][0]
}
