package main

import (
	"fmt"
)

func main(){
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

	fmt.Println("ebt",EBT)
	fmt.Println("profit",profit)
	fmt.Println("ratio",ratio)


 }