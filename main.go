package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")

	// Tipe data number
	fmt.Println("Satu = ", 1)
	fmt.Println("Dua = ", 2)
	fmt.Println("Tiga Koma Lima = ", 3.5)

	// Tipe data boolean
	fmt.Println("Benar = ", true)
	fmt.Println("Salah = ", false)

	// Tipe data string
	fmt.Println("Kata = ", "Hello")
	fmt.Println("Kata = ", "Hello, Oky Wahyu Setyaji")

	fmt.Println(len("Oky"))             // fungsi untuk mengetahui panjang string
	fmt.Println("Oky Wahyu Setyaji"[0]) // fungsi untuk mengetahui karakter pada string berdasarkan index dalam bentuk byte

	var nama string   // deklarasi variabelnya saja dan tipe datanya
	var umur int = 20 // deklarasi variabel beserta tipe data dan nilai awal
	nama = "Oky Wahyu Setyaji"
	username := "okyws" // deklarasi variabel secara singkat

	fmt.Println(nama)
	fmt.Println(umur)
	fmt.Println(username)
}
