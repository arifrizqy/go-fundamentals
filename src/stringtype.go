package main

import "fmt"

func main() {
	// Tipe data string merupakan tipe data kumpulan karakter.
	// Jumlah karakter di dalam string dapat tidak memiliki isi (tidak ada karakter sama sekali) sampai tak hingga.
	// Tipe data string di go-lang, direpresentasikan dengan kata kunci `string`.
	// Nilai data string di go-lang selalu diawali dengan karakter " (petik dua) dan diakhiri dengan karakter " (petik dua).

	fmt.Println("[String] ini sudah merupakan string")
	fmt.Println("") // String kosong

	// Beberapa contoh fungsi bawaan untuk tipe data string.
	// | Function => Keterangan |
	// len(isi string) => digunakan untuk menghitung jumlah karakter yang ada di dalam string.
	// "isi string"[number/index] => digunakan untuk mengambil karakter pada posisi/index yang ditentukan,
	// dan akan mengembalikan dalam bentuk byte bukan string karakternya. Untuk cara mengubahkan akan di bahas pada file variable.go

	fmt.Println("Panjang string:", len("ini sudah merupakan string"))
	fmt.Println("Karakter pada posisi index 5:", "ini sudah merupakan string"[5])
}
