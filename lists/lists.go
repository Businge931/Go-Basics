package lists

import "fmt"

func main() {
	prices := []float64{10.99, 8.99}
	fmt.Println(prices)

	updatedPRICES := append(prices, 5.6, 12, 25.7)
	fmt.Println(updatedPRICES)

	discountedPrices := []float64{101.99, 80.99, 20.59}
	prices= append(prices, discountedPrices...)
	fmt.Println(prices)

}

// func main() {
// 	var productNames [4]string = [4]string{"A book"}
// 	prices := [4]float64{10.99, 9.99, 20.00, 45.80}
// 	fmt.Println(prices)

// 	productNames[2] = "A carpet"
// 	// fmt.Println(productNames)

// 	fmt.Println(prices[2])
// 	featuredPrices := prices[:3]
// 	// featuredPrices = prices[1:]
// 	fmt.Println(featuredPrices)

// }
