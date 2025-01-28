package main

import "fmt"

func main() {
	// Operasi boolean dapat juga disebut dengan operasi logika.
	// Operasi logika adalah operasi untuk membandingkan dua atau lebih data/value/nilai dengan tipe data boolean.
	// Operasi ini menghasilkan sebuah nilai baru dengan tipe data boolean.
	// | Operator Logika | Keterangan |
	// (&&) => AND|dan
	// (||) => OR|atau
	// (!) => NOT|bukan

	// Operasi AND dan table kebenaran
	fmt.Println(true && true)   // => true (kedua operand bernilai true, operand1 && operand2)
	fmt.Println(true && false)  // => false
	fmt.Println(false && true)  // => false
	fmt.Println(false && false) // => false
	// Dapat disimpulkan bahwa operasi AND hanya bernilai true jika kedua operand bernilai true

	// Operasi OR dan table kebenaran
	fmt.Println(true || true)   // => true
	fmt.Println(true || false)  // => true
	fmt.Println(false || true)  // => true
	fmt.Println(false || false) // => false
	// Dapat disimpulkan bahwa operasi OR bernilai true jika salah satu operand bernilai true

	// Operasi NOT dan table kebenaran
	fmt.Println(!true)  // => false
	fmt.Println(!false) // true
	// Seperti yang telah dibahas sebelumnya pada file aritmetic.go pada pembahasan operasi unary.
	// Operasi NOT yaitu operasi untuk membalikkan nilai boolen.

	// Operasi ini juga bisa dikombinasikan dengan operasi perbandingan,
	// dikarenakan operasi perbandingan menghasilkan nilai boolean.
	var (
		nilai1 = 10
		nilai2 = 100
		nilai3 = 1000
		nilai4 = nilai1 * nilai2
	)

	fmt.Println((nilai1 < nilai2) && (nilai3 == nilai4)) // => true
	// dan seterusnya...
}
