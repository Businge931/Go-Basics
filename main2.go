package main

import "fmt"

func main() {
	userNames := make([]string, 2, 5)
	userNames = append(userNames, "Jim")

	userNames[0] = "Cathy"  //But without the make function, this would fail

	fmt.Println(userNames)
}
