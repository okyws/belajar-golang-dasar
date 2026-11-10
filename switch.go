package main

import "fmt"

func main() {
	name := "oky"

	switch name {
	case "oky":
		fmt.Println("Hello Oky")
	case "wahyu":
		fmt.Println("Hello Wahyu")
	case "setyaji":
		fmt.Println("Hello Setyaji")
	default:
		fmt.Println("Hello Guest")
	}

	// short statement switch
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	default:
		fmt.Println("Nama sudah benar")
	}

	// switch tanpa kondisi
	name = "okywahyusetyaji"
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}
