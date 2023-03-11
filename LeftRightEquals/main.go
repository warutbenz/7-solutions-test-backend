package main

import (
	"fmt"
)

func main() {
	var input string
	fmt.Print("Enter input : ")
	fmt.Scanln(&input)
	decode := decodeLeftRightEquals(input)
	fmt.Printf("output : %v\n", decode)
}

func decodeLeftRightEquals(s string) string {
	output := ""

	var numbers []int
	num := 0
	min := 0
	numbers = append(numbers, num)

	for _, char := range s {
		switch char {
		case 'L':
			num--
			numbers = append(numbers, num)
			if min > num {
				min = num
			}
		case 'R':
			num++
			numbers = append(numbers, num)
		case '=':
			numbers = append(numbers, num)
		}
	}

	if min < 0 {
		min = min * (-1)
	}

	for _, n := range numbers {
		d := n + min
		output += fmt.Sprint(d)
	}
	return output
}

func encodeLeftRightEquals(num string) string {
	result := ""
	for i := 0; i < len(num)-1; i++ {
		left := num[i]
		right := num[i+1]

		if left > right {
			result += "L"
		} else if left < right {
			result += "R"
		} else {
			result += "="
		}
	}
	return result
}
