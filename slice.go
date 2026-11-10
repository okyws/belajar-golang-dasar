package main

import "fmt"

func main() {
	names := [...]string{"oky", "wahyu", "setyaji", "putra", "putri"}
	slice1 := names[3:5]           // [putra, putri] array dari index 3 sampai sebelum index 5
	slice2 := names[1:]            // [wahyu, setyaji, putra, putri] array dari index 1 sampai akhir
	slice3 := names[:3]            // [oky, wahyu, setyaji] array dari awal sampai index sebelum 3
	var slice4 []string = names[:] // [oky, wahyu, setyaji, putra, putri] copy semua array
	fmt.Println(slice1, slice2, slice3, slice4)

	days := [...]string{"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}
	daySlice1 := days[5:] // [sabtu, minggu]
	fmt.Println(daySlice1)
	daySlice1[0] = "sabtu libur"
	daySlice1[1] = "minggu libur"
	daySlice2 := append(daySlice1, "baru liburan")
	fmt.Println(daySlice1)
	fmt.Println(daySlice2)
	fmt.Println(days)

	// membuat slice baru
	var newSlice []string = make([]string, 2, 5)
	newSlice[0] = "oky"
	newSlice[1] = "wahyu"
	// newSlice[2] = "setyaji" // error karena kapasitas slice hanya 2
	fmt.Println(newSlice)
	fmt.Println(len(newSlice)) // panjang slice
	fmt.Println(cap(newSlice)) // kapasitas slice

	newSlice2 := append(newSlice, "setyaji")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice)) // panjang slice menjadi 3 karena append
	fmt.Println(cap(newSlice)) // kapasitas slice tetap sama

	newSlice2[0] = "owl"
	fmt.Println(newSlice2)

	// mengcopy slice
	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))
	copy(toSlice, fromSlice)
	fmt.Println(toSlice, fromSlice)

	// perbedaan array dan slice
	iniArray := [...]int{1, 2, 3, 4, 5} // [5]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
