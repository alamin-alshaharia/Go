package main

import (
	"errors"
	"fmt"
)

func main() {

	title, description, err := getNoteData()

	if err != nil {
		fmt.Print(err)
		return
	}

}

func getNoteData() (string, string, error) {
	title, err := getUserInput("Enter the title:")
	if err != nil {
		fmt.Print(err)
		return "", "", nil
	}
	description, err := getUserInput("Enter the Description:")
	if err != nil {
		fmt.Print(err)
		return "", "", nil
	}
	return title, description, nil

}

func getUserInput(prompt string) (string, error) {

	fmt.Print(prompt)
	var value string

	fmt.Scan(&value)
	if value == "" {
		return "", errors.New("you must enter a vlue a vluer")

	}

	return value, nil

}
