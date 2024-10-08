package main

import "fmt"

// nil atau data kosong
// nil hanya bisa digunakan untuk tipe data tertentu seperti interface, map, slice, function, pointer dan channel

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	var person map[string]string = NewMap("Eko")

	fmt.Println(person)
}
