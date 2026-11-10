package main

import "fmt"

func main() {
	//inisialisasi map secara langsung
	person := map[string]string{
		"name": "oky",
		"age":  "20",
	}
	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["age"])

	// membuat map kosong
	book := make(map[string]string)
	book["title"] = "belajar golang"
	book["author"] = "oky"
	book["year"] = "2022"
	fmt.Println(book)
	delete(book, "author")
	fmt.Println(book)
}
