package main

import (
	"fmt"
)

// Soal Latihan Array 1

// 1. Buatlah 2 buah array berukuran 20 dan isilah dengan 20 angka acak
// 2. Buatlah program yang akan menampilkan semua anggota yang berbeda dari kedua array tersebut

func main() {

	array1 := []int{5, 6, 7, 90, 2, 3, 4, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	array2 := []int{1, 12, 13, 14, 15, 16, 17, 18, 19, 20, 10, 20, 20, 50, 50, 60, 70, 80, 90, 7}

	//Menemukan Perbedaan Dua Array:
	result := difference(array1, array2)

	fmt.Println("Array 1:", result)

	result2 := difference(array2, array1)

	fmt.Println("Array 2:", result2)

}

//Fungsi Untuk Membandingkan Dua Array
func difference(a, b []int) []int {
	//Key: Map Tipe Data: Int Value: Bool
	target := map[int]bool{}
	for _, x := range b {
		target[x] = true
	}

	result := []int{}
	for _, x := range a {
		if _, ok := target[x]; !ok {
			result = append(result, x)
		}
	}

	return result
}
