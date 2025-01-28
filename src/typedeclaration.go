package main

import "fmt"

func main() {
	// Type Declaration merupakan salah satu kemampuan milik go-lang, yang bertujuan untuk membuat tipe data 'alias'
	// dari tipe data yang sudah ada.

	type NoKTP string

	var ktpSaya NoKTP = "3215102912010005"
	fmt.Println(ktpSaya)

	fmt.Println(NoKTP("3215102912010002")) // Dapat juga digunakan untuk konversi.
}
