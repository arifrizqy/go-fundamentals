package main

import "fmt"

func main() {
	// Operasi perbandingan adalah operasi untuk membandingkan dua atau lebih data/value/nilai.
	// Operasi ini menghasilkan sebuah nilai baru dengan tipe data boolean.
	// Jika operasi dari perbandingan bernilai benar, maka hasilnya akan bernilai true.
	// Jika operasi dari perbandingan bernilai salah, maka hasilnya akan bernilai false.
	// | Operator Pembanding | Keterangan |
	// (>) => lebih besar dari
	// (<) => lebih kecil dari
	// (>=) => lebih besar atau sama dengan
	// (<=) => lebih kecil atau sama dengan
	// (==) => sama dengan
	// (!=) => tidak sama dengan

	var (
		num1 = 10
		num2 = 11
	)

	fmt.Println(num1 > num2)
	fmt.Println(num1 < num2)
	fmt.Println(num1 >= num2)
	fmt.Println(num1 <= num2)
	fmt.Println(num1 == num2)
	fmt.Println(num1 != num2)
	fmt.Println("Rizqy" == "rizqy")
	fmt.Println("Rizqy" != "rizqy")
}
