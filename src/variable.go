package main

import "fmt"

func main() {
	// Variable
	// Variable dapat diibaratkan sebagai tempat untuk menyimpan data/value/nilai.
	// Ada eberapa cara untuk mendeklarasikan variable di go-lang.

	// 1. Menggunakan kata kunci var dengan diikuti dengan nama variable dan tipe datanya,
	// kemudian mengisikan data/value/nilai nanti.
	var name string
	name = "M. Arif Rizqy F."
	fmt.Println(name)

	// 2. Menggunakan kata kunci var dengan diikuti dengan nama variable dan langsung mengisikan data/value/nilai,
	// tanpa harus menyebutkan tipe datanya.
	var nickname = "arifrizqy_." // namun kita juga dapat menyebutkan tipe datanya `var nickname string = "arifrizqy_."`,
	fmt.Println(nickname)        // akan menjadi opsional jika kita secara langsung mengisikan data/value/nilai.

	// 3. Mengganti kata kunci var dengan operator penugasan `:=`,
	// jika menggunakan metode ini kita diwajibkan untuk langsung mengisikan data/value/nilai.
	age := 20
	fmt.Println(age)

	// Dengan menggunakan beberapa metode deklarasi variable diatas, kita dapat mengganti data/value/nilai dari variable - variable yang sudah kita deklarasikan dengan cara seperti berikut.
	name = "Arif Rizqy"  // panggil nama variable yang telah dideklarasikan sebelumnya,
	nickname = "arriz_." // kemudian gunakan operator penugasan `=` untuk melakukan penyimpanan ulang data/value/nilai baru.
	age = 21             // Dalam melakukan penyimpanan ulang sebuah data/value/nilai ke variable yang sudah ada sebelumnya,

	// Constant Variable
	// Constant Variable berbeda dengan variable yang sebelumnya kita bahas (diatas).
	// Constant Variable merupakan variable yang data/value/nilai nya tidak dapat diubah.
	// Untuk mendeklarasikan constant variable di go-lang, kita dapat menggunakan kata kunci `const`.
	const PI = 3.14
	fmt.Println("Nilai PI:", PI)

	// PI = 3.15 // jika kita memaksa untuk merubah data/value/nilai yang telah kita masukkan sebelumnya,
	// pada saat pembuatan constant variable, maka akan terjadi error.

	// Multiple Variable Declaration
	// Dengan menggunakanmetode ini memungkinkan kita untuk mendeklarasikan beberapa variable sekaligus.
	var (
		username = "arifrizqy"
		password = "123"
	) // Jika menggunakan kata kunci `var`, maka data/value/nilai nya dapat diubah.

	const (
		daysInWeek    = 7
		minutesInHour = 60
	) // Jika menggunakan kata kunci `const`, maka data/value/nilai nya tidak dapat diubah.

	fmt.Println(username, password)
	fmt.Println(daysInWeek, minutesInHour)
}
