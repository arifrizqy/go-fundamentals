package main

import "fmt"

func main() {
	// Aritmatika adalah sebuah operasi matematika pada umumnya.
	// | Operator | Keterangan |
	// (+) | Penjumlahan
	// (-) | Pengurangan
	// (*) | Perkalian
	// (/) | Pembagian
	// (%) | Sisa Pembagian

	fmt.Println(1 + 24) // Penjumlahan
	fmt.Println(25 - 1) // Pengurangan
	fmt.Println(4 * 4)  // Perkalian
	fmt.Println(24 / 6) // Pembagian
	fmt.Println(24 % 4) // Sisa Pembagian

	var (
		num1 = 10
		num2 = 7
		num3 = (num1-num2)*2 - num1
	)

	fmt.Println(num1, num2, num3)

	// Augmented Assignment
	// Dalam aritmatika terdapat juga yang bernama augmented assignment.
	// Augmented assignment dapat saya sebut sebagai operasi kediri sendiri.
	// | Operator Augmented Assignment | Operasi Matematika |
	// (+=) | a += 10 => a = a + 10
	// (-=) | a -= 2.5 => a = a - 2.5
	// (*=) | a *= 3.14 => a = a * 3.14
	// (/=) | a /= 1 => a = a / 1
	// (%=) | a %= 2 => a = a % 2

	var angka1 int = 10
	angka2 := 5

	angka1 += angka2    // angka1 = angka1 + angka2. Penggunaan operator augmented assignment hanya dapat digunakan
	fmt.Println(angka1) // pada variable yang telah memiliki data/value/nilai. Jadi tidak akan bisa digunakan pada saat awal pendeklarasian variable.

	// Unary Operator
	// Jika kalian ingin melakukan operasi matematika misal (a += 1), hanya bertambah sebanyak 1 saja.
	// Maka lebih baik gunakan unary operator.
	// | Operator Unary | Keterangan |
	// (++) | a++ => a = a + 1
	// (--) | b-- => b = b - 1

	a := 10
	a++
	fmt.Println(a)

	// Dalam numeric biasanya bilangan dapat berupa negative atau positive.
	// Gunakan operator unary untuk membuat sebuah bilangan menjadi negative atau positive.
	// | Operator Unary | Keterangan |
	// (-) | -a => -10
	// (+) | +b => +314 || 314 // Pada dasarnya jika kita menuliskan sebuah angka tanpa tanda (-) maka angka tersebut akan menjadi positive.

	fmt.Println(-50)

	// Dalam unary operator juga terdapat operator untuk digunakan untuk tipe data boolean.
	// | Operator Unary | Keterangan |
	// (!) | !true => false // Operasi ini dapat disebut sebagai negasi, atau lebih gampangnya disebut kebalikan.

	fmt.Println(!true)
}
