package main

import "fmt"

func main() {
	fmt.Println(Max([]int{1, 2, 4, 66, 323, 1, 52, 521, 234}))
}

func Max(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	var max int = numbers[0]

	for _, number := range numbers {
		if number > max {
			max = number
		}
	}
	return max
}

func MaxIndex(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	var max int = numbers[0]
	var index int

	for i, number := range numbers {
		if number > max {
			max = number
			index = i
		}

	}

	return index
}
