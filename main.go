package main

import (
	"errors"
	"fmt"
)

type str string

func main() {
	title, content, err := getNoteData()

	if err != nil {
		fmt.Println(err)
		return
	}
}

func getNoteData() (str, str, error) {
	title, err := getUserInput("Note title: ")

	if err != nil {
		return "", "", err
	}

	content, err := getUserInput("Note Content")

	if err != nil {
		return "", "", err
	}

	return title, content, nil
}

func getUserInput(prompt string) (str, error) {
	fmt.Print(prompt)
	var value str
	fmt.Scan(&value)

	if value == "" {
		return "", errors.New("invalide input")
	}
	return value, nil
}
