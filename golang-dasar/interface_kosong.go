package main

import "fmt"

func Ups(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "Ups"
	}
}

// anggap saja bahwa interface kosong tipe data any, jadi jika digunakan sebagai return value, tipe data apa saja boleh digunakan
func main() {
	var data interface{} = Ups(2)
	fmt.Println(data)
}
