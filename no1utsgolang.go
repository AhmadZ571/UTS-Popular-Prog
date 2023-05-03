package main

import "fmt"

func sum(numbers ...int) int {
	total := 0
	for _, n := range numbers {
		total += n
	}
	return total
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	result := sum(nums...)
	fmt.Println("The sum of", nums, "is", result)
}
