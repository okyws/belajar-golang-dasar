package main

import "fmt"

func main() {
	name := "Oky"
	if name == "Wahyu" {
		fmt.Println("Hello Wahyu")
	} else if name == "Oky" {
		fmt.Println("Hello Oky")
	} else if name == "Setyaji" {
		fmt.Println("Hello Setyaji")
	} else {
		fmt.Println("Hello Guest")
	}

	// if dengan short statement
	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
