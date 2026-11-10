package main

import "fmt"

func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Oky"
	middleName = "Wahyu"
	lastName = "Setyaji"

	return firstName, middleName, lastName
}

func main() {
	a, b, c := getCompleteName()
	fmt.Println(a, b, c)
}
