package main

// import (
// 	"fmt"
// 	"math"
// )
// const inflationRate =2.5

// func main(){
// 	var investmentAmount float64 
// 	var years float64 
// 	var expectedReturnRate float64 = 5.5

// 	// fmt.Print("Investment Amount: ")
// 	outputText("Investment Amount: ")
// 	fmt.Scan(&investmentAmount)

// 	// fmt.Print("years: ")
// 	outputText("years: ")
// 	fmt.Scan(&years)

// 	// fmt.Print("Expected Return: ")
// 	outputText("Expected Return: ")
// 	fmt.Scan(&expectedReturnRate)

// 	futureValue,futureRealValue := calculateFutureValues(investmentAmount,expectedReturnRate,years)
//     // futureValue := investmentAmount * math.Pow(1 + expectedReturnRate /100,years) 
// 	// futureRealValue := futureValue / math.Pow(1 + inflationRate /100,years)

// 	fmt.Print(futureRealValue,futureValue)
//  }

// func outputText(text string){
// 	fmt.Print(text)
// } 

// func calculateFutureValues(investmentAmount ,expectedReturnRate, years float64) (fv float64,rfv float64){
// 	fv =investmentAmount * math.Pow(1 + expectedReturnRate /100,years)
// 	rfv = fv / math.Pow(1 + inflationRate /100,years)
// 	return fv,  rfv
// 	// return
// }