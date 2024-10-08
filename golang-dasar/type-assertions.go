// kemampuan untuk merubaha tipe data menjadi tipe data yang diinginkan
// sering terjadi pada perubahan tipe data interface kosong (any) ke tipe data yang diinginkan

package main

import "fmt"

func random() any {
	return "OK"
}

func main() {
	var result any = random()
	var resultString string = result.(string)
	fmt.Println(resultString)

	// jadi terlebih dahulu cek tipe data hasil return agar tidak terjadi panic saat menkonfersi tipe data
	// pengecekan tipe data bisa menggunakan fmt.Printf  dengan simbol persen T di dalamnya
	fmt.Printf("result = %T", result)
}
