package main

import "fmt"

func main() {
	numbers := []int{3, 6, 7, 9, 5, 2}
	total := sumUp(3, 6, 7, 9, 5, 2)
	anotherSum := sumUp(1, numbers...)

	fmt.Println(total)
	fmt.Println(anotherSum)
}

func sumUp(startingValue int, numbers ...int) int {
	sum := 0

	for _, val := range numbers {
		sum += val
	}
	return sum
}
