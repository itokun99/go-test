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

	return x
}

func Max(x, y int) int {

	return x
}

/*
2. lengkapi program berikut untuk reverse urutan dari array bilangan integer
	 	contoh:
				input {1,2,3,4,5}
				ouput {5,4,3,2,1}
*/
func reverse(sw []int) {

}

func main() {
	//jawaban 1
	x := []int{1, 2, 3, 4, 5}
	reverse(x)
	fmt.Println(x)

	//jawaban 2
	fmt.Println("MIN:", Min(4, 5))
	fmt.Println("MAX:", Max(4, 5))

}
