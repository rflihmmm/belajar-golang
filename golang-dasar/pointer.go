package main

import "fmt"

// secara default di golang semua variable itu passing by value, atau menduplikasi isi data dari varibale lain
// jadi jika terjadi perubahan pada variable yang di ambil datanya, varibale yang mengambil data tersebut tidak ikut berubah

type Address struct {
	City, Province, Country string
}

func main() {
	// contoh pass by value
	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	address2 := address1

	address1.City = "Bandung"
	fmt.Println(address1)
	fmt.Println(address2)

	// contoh pointer atau pass by reference
	// pointer di identifikasikan dengan simbol "&" di depan nama variable yang ingin di pass by reference
	address3 := &address1
	var address4 *Address = &address2
	// kedua penulisan di atas sama, yang merujuk pada pointer
	address3.City = "Cirebon"
	// jadi jika address3 yang datanya merujuk pada variable address1 di ubah, maka address1 juga akan ikut berubah sesuai dengan perubahan yang terjadi pada address3

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)
	fmt.Println(address4)
}
