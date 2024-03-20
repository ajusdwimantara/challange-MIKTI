package main

import (
	"fmt"
)

// struct untuk menyimpan massa dan tinggi
/*
	menggunakan pointer agar lebih hemat memory (tidak melakukan copy)
*/
type PersonInfo struct {
	massa, tinggi *float32
}

// fungsi untuk mengetahui apakah BMI mark lebih tinggi
// return true jika benar dan false jika salah
func isMarkHigherBMI(mark *PersonInfo, john *PersonInfo) bool { // assign by reference agar lebih hemat memory (tidak melakukan copy)
	mark_bmi := *mark.massa / ((*mark.tinggi) * (*mark.tinggi)) // hitung bmi mark
	john_bmi := *john.massa / ((*john.tinggi) * (*john.tinggi)) // hitung bmi john

	fmt.Printf("Mark BMI : %.2f\n", mark_bmi)
	fmt.Printf("John BMI : %.2f\n", john_bmi)

	if mark_bmi > john_bmi {
		return true
	}
	return false
}
func main() {
	// initiate variable untuk data dan apakah BMI Mark lebih tinggi dari John
	var berat_mark float32
	var tinggi_mark float32

	var berat_john float32
	var tinggi_john float32

	var markHigherBMI bool

	// ======== dataset 1 ======//
	fmt.Println("DATASET 1")
	berat_mark = 78
	tinggi_mark = 1.69

	berat_john = 92
	tinggi_john = 1.95

	markHigherBMI = false

	// assign data ke struct
	mark := PersonInfo{&berat_mark, &tinggi_mark}
	john := PersonInfo{&berat_john, &tinggi_john}

	markHigherBMI = isMarkHigherBMI(&mark, &john)

	if markHigherBMI {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
	fmt.Println()

	// ======== dataset 2 ======//
	fmt.Println("DATASET 2")
	// ubah data yang di-point
	berat_mark = 95
	tinggi_mark = 1.88

	berat_john = 85
	tinggi_john = 1.76

	markHigherBMI = false

	markHigherBMI = isMarkHigherBMI(&mark, &john)

	if markHigherBMI {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
	fmt.Println()

}
