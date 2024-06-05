package main

import "fmt"

type floatMap map[string]float64

func(m floatMap) output(){
	fmt.Println(m)
}

func main() {
	userNames := make([]string, 2, 5)
	userNames = append(userNames, "Jim")

	userNames[0] = "Cathy" //But without the make function, this would fail

	fmt.Println(userNames)

	courseRatings := make(floatMap,3)

	courseRatings["go"]=4.7
	courseRatings["react"]=4.8
	courseRatings["node"]=4.8

	// fmt.Println(courseRatings)
	courseRatings.output()

}
