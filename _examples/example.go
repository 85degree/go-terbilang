package main

import (
	"fmt"

	"github.com/85degree/go-terbilang"
)

func main() {
	// fungsi terbilang ini dapat terdiri dari 3 fungsi utama
	var angka int
	angka = 109209
	// fungsi ToTerbilang(angka) menampilkan terbilang dalam format huruf
	// dari angka 109209 menjadi seratus sembilan ribu dua ratus sembilan
	fmt.Println(terbilang.ToTerbilang(angka))

	// fungsi berikutnya yaitu dengan ToTerbilangRp(angka)
	// fungsi ini menambahkan mata uang rupiah ke dalam hasil terbilang
	fmt.Println(terbilang.ToTerbilangRp(angka))

	// fungsi yang terakhir adalah ToTerbilangSuffix(angka)
	// fungsi ini mirip dengan ToTerbilangRp(angka) namun
	// kita dapat mengubah rupiah menjadi apapun yang kita inginkan
	// seperti: dollar, bata, buah, ekor, dll
	fmt.Println(terbilang.ToTerbilangSuffix("ekor", 102))
}
