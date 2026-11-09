package main

import "fmt"

func main() {
	var name [3]string
	name[0] = "Oky"
	name[1] = "Wahyu"
	name[2] = "Setyaji"
	fmt.Println(name[0])
	fmt.Println(name[1])
	fmt.Println(name[2])

	// array literal, untuk bisa menggunakan ..., array harus langsung diisi
	var value = [...]int{
		10,
		20,
		30,
	}

	fmt.Println(len(value)) // panjang array

	value[1] = 100 // mengubah nilai array index ke 1

	fmt.Println(value[0])
	fmt.Println(value[1])
	fmt.Println(value[2])
}
