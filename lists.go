package main

import "fmt"

type Product struct {
	id    string
	title string
	price float64
}

func main() {
	// prices := []float64{10.99, 8.99}
	// fmt.Println(prices)

	// updatedPRICES:= append(prices,5.6)
	// fmt.Println(updatedPRICES)

	//**************** EXERCISE SOLUTION **************

	//1)
	hobbies := [3]string{"Chilling", "Watching series", "Good company"}
	fmt.Println(hobbies)

	//2)
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:3])

	//3)
	mainHobbies := hobbies[0:2]
	fmt.Println(mainHobbies)

	//4)
	fmt.Println(cap(mainHobbies))
	mainHobbies = mainHobbies[1:3]
	fmt.Println(mainHobbies)

	//5)
	courseGoals := []string{"Learn Go", "Learn all the basics"}
	fmt.Println(courseGoals)

	//6)
	courseGoals[1] = "Learn all details"
	courseGoals = append(courseGoals, "Master everything")
	fmt.Println(courseGoals)

	//7)
	products := []Product{
		{
			"first-product",
			"A first product",
			12.99,
		},
		{
			"secont-product",
			"A secont product",
			12.99,
		},
	}
	fmt.Println(products)

	newProduct := Product{
		"third-product", "Third product", 15.99,
	}
	products =append(products, newProduct)
	fmt.Println(products)

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
