package main

import "fmt"

func main() {
	// Array atau dalam bahasa disebut larik.
	// Array merupakan tipe data yang berisikan kumpulan data dengan tipe data yang sama.
	// Saat membuat array, kita perlu menentukan jumlah data yang dapat ditampung oleh array.
	// Daya tampung array tidak dapat berubah (berkurang maupun bertambah) lebih dari / kurang dari daya tampung yang telah ditentukan.

	// Dalam array terdapat yang namanya index.
	// Index merupakan posisi dari suatu data di dalam array.
	// Jika kita ingin mengakses sebuah data di dalam array, maka kita harus menggunakan index untuk mendapatkan data tersebut.
	// Data pertama dalam array berindex 0, data kedua berindex 1, dan seterusnya.

	// Cara membuat array
	// 1
	var names [3]string // string menandakan tipe data yang dapat ditampung oleh array hanya data/value/nilai bertipe data string saja.
	names[0] = "arif"
	names[1] = "rizqy"
	names[2] = "developer"
	// names[3] = "arriz" // akan error, dikarenakan daya tampung telah kita tentukan sebelumnya yaitu 3.

	fmt.Println(names, names[0], names[1], names[2])

	// 2
	var nilai = [3]int{90, 80, 85}

	fmt.Println(nilai, nilai[0], nilai[1], nilai[2])

	// 3
	var (
		nama_mhs  = [3]string{"arif", "eko", "forsaken"}
		nilai_mhs = [...]int{90, 80, 77} // tanda ... digunakan jika kita sebelumnya belum yakin nantinya berapa jumlah data yang dapat ditampung oleh array.
	)

	fmt.Println(nama_mhs[0], "memiliki nilai", nilai_mhs[0])
	fmt.Println(nama_mhs[1], "memiliki nilai", nilai_mhs[1])
	fmt.Println(nama_mhs[2], "memiliki nilai", nilai_mhs[2])
	// masih terdapat beberapa cara untuk mendeklarasikan array. untuk selebihnya silahkan explore lebih lanjut.

	// Function yang dapat digunakan untuk array
	// | Function | Keterangan |
	// len(array) => digunakan untuk mendapatkan daya tampung dari sebuah array.
	// array[index] => digunakan untuk mengakses atau mengambil data yang berada di dalam array.
	// array[index] = value => igunakan untuk mengubah data/value/nilai yang ada di dalam array dengan data/value/nilai yang baru.
	fmt.Println("Panjang array:", len(names))

	nilai_mhs[2] = 85
	fmt.Println("Nilai", nama_mhs[2], "berubah menjadi", nilai_mhs[2])
}
