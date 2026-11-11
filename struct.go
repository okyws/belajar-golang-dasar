package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "my name is", customer.Name)
}

func main() {
	var oky Customer
	fmt.Println(oky)

	oky.Name = "Oky Wahyu Setyaji"
	oky.Address = "Indonesia"
	oky.Age = 28
	fmt.Println(oky)
	fmt.Println(oky.Name)
	fmt.Println(oky.Address)
	fmt.Println(oky.Age)

	joko := Customer{
		Name:    "Joko",
		Address: "Indonesia",
		Age:     30,
	}
	fmt.Println(joko)

	budi := Customer{"Budi", "Indonesia", 30}
	fmt.Println(budi)

	budi.sayHello("Agus")
	oky.sayHello("Agus")
	joko.sayHello("Agus")
}
