# Belajar Golang Dasar

Kumpulan contoh program Go untuk belajar dasar-dasar bahasa pemrograman Go. Materi disusun sebagai file-file kecil agar konsepnya mudah dibaca dan dicoba satu per satu. Project ini mengikuti pembelajaran Golang Dasar dari **Programmer Zaman Now (PZN)**.

## Persiapan

Pastikan Go sudah terpasang. Periksa instalasi dengan:

```bash
go version
```

## Menjalankan contoh

Setiap file contoh mendeklarasikan `package main` dan fungsi `main` masing-masing. Karena itu, jalankan satu file contoh dalam satu waktu dari direktori project:

```bash
go run array.go
```

Ganti `array.go` dengan nama file yang ingin dipelajari, misalnya:

```bash
go run main.go
go run conversion.go
go run operasi-boolean.go
```

> Hindari menjalankan `go run .` untuk kumpulan file ini sekaligus: beberapa file memiliki fungsi `main`, sehingga akan terjadi deklarasi fungsi ganda.

## Menjalankan dari awal

1. Clone atau unduh repository ini.
2. Buka terminal pada direktori project.
3. Jalankan salah satu file dengan `go run nama-file.go`.
4. Baca dan ubah contoh, lalu jalankan kembali untuk melihat hasilnya.

## Modul

Project ini menggunakan Go Modules. Informasi modul terdapat di `go.mod`.
