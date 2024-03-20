package main

import "fmt"

// struct untuk menyimpan poin tim lumba2
type Team struct {
	/*
		menggunakan pointer agar lebih hemat memory (tidak melakukan copy)
	*/
	poin [3]*int // int sesuai data pada dataset
}

// fungsi untuk membandingkan rata2 kedua tim sekaligus menentukan pemenang
func compare(lumba2 *Team, koala *Team) { // assign by reference agar lebih hemat memory (tidak melakukan copy)
	// hitung rata2 pada tiap tim
	n := 0
	avg_lumba2 := 0
	avg_koala := 0

	for i := 0; i < len(lumba2.poin); i++ {
		n += 1
		avg_lumba2 += *lumba2.poin[i]
		avg_koala += *koala.poin[i]
	}
	avg_lumba2 /= n
	avg_koala /= n

	fmt.Printf("Skor rata-rata tim lumba-lumba: %d\n", avg_lumba2)
	fmt.Printf("Skor rata-rata tim koala: %d\n", avg_koala)

	// conditional statement sesuai no 3 dan 4
	if avg_koala >= 100 || avg_lumba2 >= 100 {
		if avg_koala > avg_lumba2 {
			fmt.Println("Koala memenangkan trofi!")
		} else if avg_lumba2 > avg_koala {
			fmt.Println("Lumba-lumba memenangkan trofi!")
		} else {
			fmt.Println("Seri!")
		}
	} else {
		fmt.Println("Tidak ada tim yang memenangkan trofi!")
	}
}

func main() {
	// variable untuk menyimpan data
	poin_lumba2 := [3]int{
		0, 0, 0,
	}

	poin_koala := [3]int{
		0, 0, 0,
	}

	// ======== dataset 1 ======== //
	poin_lumba2[0] = 96
	poin_lumba2[1] = 108
	poin_lumba2[2] = 89

	poin_koala[0] = 88
	poin_koala[1] = 91
	poin_koala[2] = 110

	// assign data ke struct
	lumba2 := Team{[3]*int{
		&poin_lumba2[0], &poin_lumba2[1], &poin_lumba2[2],
	}}
	koala := Team{[3]*int{
		&poin_koala[0], &poin_koala[1], &poin_koala[2],
	}}

	// compare untuk mencari tim yang menang
	fmt.Println("DATASET 1")
	compare(&lumba2, &koala)
	fmt.Println()

	// ======== dataset 2 ======= //
	// ubah data yang di-point
	poin_lumba2[0] = 97
	poin_lumba2[1] = 112
	poin_lumba2[2] = 101

	poin_koala[0] = 109
	poin_koala[1] = 95
	poin_koala[2] = 123

	// compare untuk mencari tim yang menang
	fmt.Println("DATASET 2")
	compare(&lumba2, &koala)
	fmt.Println()

	// ======== dataset 3 ======= //
	// ubah data yang di-point
	poin_lumba2[0] = 97
	poin_lumba2[1] = 112
	poin_lumba2[2] = 101

	poin_koala[0] = 109
	poin_koala[1] = 95
	poin_koala[2] = 106

	// compare untuk mencari tim yang menang
	fmt.Println("DATASET 3")
	compare(&lumba2, &koala)
	fmt.Println()
}
