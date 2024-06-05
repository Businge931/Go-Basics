package main

import "fmt"

type Product struct{}

func main() {
	var productNames [4]string = [4]string{"A book"}
	prices := [4]float64{10.99, 9.99, 20.00, 45.80}
	fmt.Println(prices)

	productNames[2] = "A carpet"
	// fmt.Println(productNames)

	fmt.Println(prices[2])
	featuredPrices := prices[1:3]
	fmt.Println(featuredPrices)


}
