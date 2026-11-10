package main

import "fmt"

func getHello(name string) string {
	hello := "Hello " + name
	return hello
}

func main() {
	result := getHello("Oky")
	fmt.Println(result)

	fmt.Println(getHello("Wahyu"))
	fmt.Println(getHello("Setyaji"))
}
