package main

import (
	"fmt"
)

func main(){
	// const inflationRate =2.5
	// var investmentAmount float64 
	// var years float64
	// var expectedReturnRate float64

	var taxRate float64
	var revenue float64
	var expenses float64

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Tax rate: ")
	fmt.Scan(&taxRate)

	EBT := revenue - expenses
    profit := EBT * (1-taxRate / 100 )
	ratio := EBT / profit

	fmt.Println(EBT)
	fmt.Println(profit)
	fmt.Println(ratio)

	// fmt.Print("years: ")
	// fmt.Scan(&years)

	// fmt.Print("Expected Return: ")
	// fmt.Scan(&expectedReturnRate)


    // futureValue := investmentAmount * math.Pow(1 + expectedReturnRate /100,years) 
	// futureRealValue :=futureValue / math.Pow(1 + inflationRate /100,years)

    // fmt.Println(futureValue)
    // fmt.Println(futureRealValue)
 }