package main

import "fmt"

func main() {
	// Dalam go-lang tipe data numeric dibagi 2, yaitu:
	// 1. Integer
	// 2. Floating Point

	// Integer -------------------------------------
	// signed integer ------------------------------
	// | Tipe Data => Nilai min. <-> Nilai max. |
	// int8 => -128 <-> 127
	// int16 => -32768 <-> 32767
	// int32 => -2147483648 <-> 2147483647
	// int64 => -9223372036854775808 <-> 9223372036854775807

	// unsigned integer ----------------------------
	// | Tipe Data => Nilai min. <-> Nilai max. |
	// uint8 => 0 <-> 255
	// uint16 => 0 <-> 65535
	// uint32 => 0 <-> 4294967295
	// uint64 => 0 <-> 18446744073709551615

	fmt.Println("[Integer] satu:", 1)
	fmt.Println("[Integer] negatif seratus dua puluh delapan:", -128)

	// Floating Point ------------------------------
	// | Tipe Data => Nilai min. <-> Nilai max. |
	// float32 => -3.4028234663852886e+38 <-> 3.4028234663852886e+38
	// float64 => -1.7976931348623157e+308 <-> 1.7976931348623157e+308

	fmt.Println("[Float] tiga koma empat belas:", 3.14)
	fmt.Println("[Float] negatif lima puluh koma lima puluh lima:", -50.55)

	// Alias dari tipe data numeric di go-lang -----
	// | Tipe data => Alias untuk |
	// byte => uint8
	// rune => int32
	// int => Min. int32
	// uint => MIn. uint32
}
