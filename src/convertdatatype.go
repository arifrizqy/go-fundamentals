package main

import "fmt"

func main() {
	// Dalam go-lang kita dapat melakukan konversi dari satu tipe data ke tipe data yang lain.
	// Misal dari tipe data int32 ke tipe data int64, dari tipe data int32 ke tipe data int12, dan sebagainya.
	// Namun untuk melakukan konversi tipe data, kita juga perlu memperhatikan beberapa ketentuan juga.

	// Pada konversi tipe data khususnya integer, melakukan konversi tipe data ke yang lebih tinggi umumnya
	// tidak akan menjadi masalah/ perubahan data/value/nilai.

	var (
		nilai32 int32 = 32768
		nilai64 int64 = int64(nilai32) // Konversi tidak akan berdampak terhadap data/value/nilai,
		// dikarenakan rentang angka yang dapat ditampungan tipe data int64 lebih besar dari pada int32, dapat di lihat di file numerictype.go
		nilai16 int16 = int16(nilai32) // Konversi ini akan menyebabkan perubahan terhadap data/value/nilai,
		// dikarenakan rentang angka yang dapat ditampungan tipe data int16 lebih kecil dari pada int32. Ini biasa dinakan number overflow.
	)

	fmt.Println(nilai32, nilai64, nilai16)

	// Pada file stringtype.go kita pernah mengetahui/ menyinggung cara untuk mendapatkan karakter dari sebuah string.
	// Akan tetapi, hasil kembalian dari cara tersebut tidak sesuai dengan keinginan kita yang menginginkan karakter tersebut,
	// melainkan mengembalikan angka dengan tipe data byte. Kita dapat melakukan konversi tipe data `byte` ke tipe data string
	// untuk mengetahui karakter apa yang telah terambil.

	var name = "arriz_."
	charFromName := name[2]
	var char string = string(charFromName)

	fmt.Println(name, charFromName, char)
}
