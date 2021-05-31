package main

import "fmt"

//Pilih salah satu soal
//Kerjakan dengan menggunakan bahasa golang

/*
1. lengkapi program berikut untuk mencari bilangan Min dan Max dari 2 bilangan input
	 	contoh:
				input 3 dan 4
				output Min = 3 Max = 4
*/
func Min(x, y int) int {

	// deklarasi variabel "result"
	var result int

	// jika x lebih kecil dari y
	if x < y {
		result = x // assign x ke result

		// jika y lebih kecil dari x
	} else if y < x {
		result = y // assign x ke result
	}

	// balikan nilai result
	return result
}

func Max(x, y int) int {

	// deklarasi variabel "result"
	var result int

	// jika x lebih besar y
	if x > y {
		result = x // assign x ke result

		// jika y lebih besar dari x
	} else if y > x {
		result = y // assign y ke result
	}

	// kembalikan nilai result
	return result
}

/*
2. lengkapi program berikut untuk reverse urutan dari array bilangan integer
	 	contoh:
				input {1,2,3,4,5}
				ouput {5,4,3,2,1}
*/
func reverse(sw []int) []int {

	// jalankan perulangan
	// deklarasikan i sebagai nilai 0
	// lakukan perulangan jika nilai i lebih kecil dari jumlah anggota array
	// nilai i akan bertambah + 1 setiap perulangan
	// for i := 0; i < len(sw)/2; i++ {
	// 	// deklarasikan nilai j dengan kalkulasi, jumlah anggota array dikurangi nilai i dan 1
	// 	j := len(sw) - i - 1

	// 	// tukar nilai array
	// 	sw[i], sw[j] = sw[j], sw[i]
	// }

	// // balikan nilai array sw
	return sw
}

func main() {
	//jawaban 1
	x := []int{1, 2, 3, 4, 5}
	reverse(x)
	fmt.Println(x)

	//jawaban 2
	fmt.Println("MIN:", Min(3, 4))
	fmt.Println("MAX:", Max(4, 5))

}
