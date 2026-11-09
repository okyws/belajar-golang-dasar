package main

import "fmt"

func main() {
	var a int = 10
	var b int = 20
	var c int = a + b
	var d int = a - b
	var e int = a * b
	var f int = b / a
	fmt.Println(a, b, c, d, e, f)

	var i int
	i = 10
	i++
	i++
	fmt.Println(i)
	i--
	fmt.Println(i)
	i += 5
	fmt.Println(i)

}
