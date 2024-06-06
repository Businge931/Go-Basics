package main

// import "fmt"

// type transformFn func(int) int

// func main() {
// 	numbers := []int{1, 2, 3, 4}
// 	moreNumbers :=[]int{5,1,2,6}
// 	doubled := transformNumbers(&numbers, double)
// 	trippled := transformNumbers(&numbers, tripple)

// 	fmt.Println(doubled)
// 	fmt.Println(trippled)

// 	transformerFn1 := getTransformerFunction(&numbers)
// 	transformerFn2 := getTransformerFunction(&moreNumbers)

// 	transformedNumbers :=transformNumbers(&numbers,transformerFn1)
// 	moreTransformedNumbers :=transformNumbers(&numbers,transformerFn2)

// 	fmt.Println(transformedNumbers)
// 	fmt.Println(moreTransformedNumbers)
// }

// func transformNumbers(numbers *[]int, transform transformFn) []int {
// 	dNumbers := []int{}
// 	for _, val := range *numbers {
// 		dNumbers = append(dNumbers, transform(val))
// 	}
// 	return dNumbers
// }

// func getTransformerFunction(numbers *[]int) transformFn {
// 	if (*numbers)[0]==1{
// 		return double
// 	}else {
// 		return tripple
// 	}
// }

// func double(number int) int {
// 	return number * 2
// }
// func tripple(number int) int {
// 	return number * 3
// }
