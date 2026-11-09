package main

import "fmt"

func main() {
	type NoKTP string

	var ktpOky NoKTP = "123456789"
	var sample = "1111111111"
	var sampleKTP NoKTP = NoKTP(sample)
	fmt.Println(ktpOky)
	fmt.Println(sampleKTP)

}
