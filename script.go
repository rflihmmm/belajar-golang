package main

func main() {
	/*
		name := "Rafli Hamdani"
		name = "Muh Rafli Hamdani"
		fmt.Println(name)

		age := 20
		fmt.Println(age)

		nim := 42221053
		fmt.Println("NIM", nim)
	*/

	/*
		var (
			firstName = "Rafli"
			lastName  = "Hamdani"
			umur      = 20
		)

		fmt.Println(firstName)
		fmt.Println(lastName)
		fmt.Println(umur)
	*/

	//var (
	//name = "Rafli"
	//e = 40
	//eString = string(e)
	//flt = 3/4
	//bulat = float32(flt)
	//)

	//fmt.Println(name)
	//fmt.Println(e)
	//fmt.Println(bulat)

	// var name [2] string
	//name[0] = "Rafli"
	// name[1]	="Hamdani"

	// fmt.Println(name)
	// fmt.Println(len(name))
	// fmt.Println(len(name[1]))

	// arr := [...]string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}

	// fmt.Println(len(arr))

	// nama := "Agus"
	// leng := nama[1]

	// fmt.Println(string(leng))

	// -------- tipe data slice -----------
	//potongan data array. slice mengacu pada array, jadi jika array diubah, maka slice pun akan berubah. Begitupun sebaliknya, jika slice diubah, array pada slice tersebut akan berubah pula.

	//bulan := [...]string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}

	// slice1 akan mengambil array bulan index ke 4 sampai sebelum index ke 7. jika bulan[4:], berarti dari index ke 4 samapi index terakhir. dan jika bulan[:7], berarti dari index paling awal samapai sebelum index ke 7.
	//slice1 := bulan[4:7]
	//fmt.Println(slice1) // ["May", "June", "July"]
	//slice1[1] = "Mei" // ["May", "Mei", "July"], array bulan juga akan beruba
	//fmt.Println(len(slice1)) // 3
	//fmt.Println(cap(slice1)) // 8 ---- cap berarti kapasitas slice tersebut. maksudnya panjang slice yang bisa didapatkan yang mengacu pada array.

	//slice2 := bulan[10:]
	//slice3 := append(slice2, "duarsember") // membuat slice baru yang datanya diambil dari slice 2, parameter kedua adalah data yang ingin ditambahkan di index paling akhir.
	//fmt.Println(slice3) // slice3 tidak akan merubah array utama dan akan membuat array baru. itu dikarenakan slice3 melebihi kapasitas dari slice2.

	// membuat slice baru yang akan membuat array baru pula dan tidak tergantung pada array yang telah di deklarasi
	// slice4 := make([]string, 2, 5) // tipedata, length, capacity
	// slice4[0] = "Rafli"
	// slice4[1] = "Hamdani"
	// //slice4[2] = "Error" // Error, karena melebihi length yang telah ditentukan yaitu 2

	// fmt.Println(slice4)
	// fmt.Println(len(slice4)) //2
	// fmt.Println(cap(slice4)) //5

	// toSlice := make([]string, len(slice4), cap(slice4))

	// copy(toSlice, slice4)

	// fmt.Println(toSlice)

	// ----------- Tipe Data map ------------
	// tipe data map pada go seperti tipe data object pada javascript. Didalamnya terdapat key dan value
	//map[tipedatakey]tipedatavalue{}
	// person := map[string]string{
	// 	"name":    "Rafli",
	// 	"address": "Pangkep",
	// }

	// person["tittle"] = "Mahasiswa" // mengubah(jika key-nya sudah ada) atau menambah data baru pada map

	// fmt.Println(person)
	// fmt.Println(person["name"])

	// // membuat varible map baru yang isinya kosong
	// book := make(map[string]string)
	// // menambah key dan value pada variable book
	// book["tittle"] = "Lembaran Berdarah Sejarah Indonesia"
	// book["isbn"] = "978-234-4546-0-1"
	// book["orang"] = "error"

	// delete(book, "orang") // menghapus key "orang" pada varible book yang bertipe data map

	// fmt.Println(book)

	// --------------- If expression ------------
	// nama := "rafli"

	// if nama == "Rafli" || nama == "rafli" {
	// 	fmt.Println("Pendek")
	// } else if nama == "Rafli Hamdani" {
	// 	fmt.Println("Panjang")
	// } else {
	// 	fmt.Println("False")
	// }

	//If Short Statement
	// if length := len(nama); length >= 5 {
	// 	fmt.Println("true")
	// }

	// --------------- Switch Expression --------------------
	// nama := "rafli"

	// switch nama {
	// case "Rafli":
	// 	fmt.Println("CamelCase")
	// case "rafli":
	// 	fmt.Println("SmallCase")
	// default:
	// 	fmt.Println("False")
	// }

	// Switch Short Statement
	// switch length := len(nama); length >= 5 {
	// case true:
	// 	fmt.Println(true)
	// case false:
	// 	fmt.Println(false)
	// }

	// Switch case tanda kondisi
	// length := len("nama")
	// switch {
	// case length > 5:
	// 	fmt.Println("Panjang")
	// case length < 5:
	// 	fmt.Println("Pendek")
	// default:
	// 	fmt.Println("Benar")
	// }

	// --------------- For Loops --------------------
	// nama := "takut untuk memulai"
	// for i := 0; i < len(nama); i++ {
	// 	if i >= len(nama) {
	// 		fmt.Print(nama[i])
	// 		fmt.Println(".")
	// 	} else {
	// 		fmt.Print(nama[i])
	// 		fmt.Print(".")
	// 	}

	// }

	// slice := []string{"Rafli", "Hamdani", "Rusmono"}

	// for i, v := range slice {
	// 	fmt.Println(i, " ", v)
	// }

	// --------------- Break and Continue ------------

	// saat menuslikan break, semua perulangan akan dihentikan saat itu juga, jadi fungsi dibawah break tidak akan di jalankan
	// for i := 0; i < 10; i++ {
	// 	if i == 5 {
	// 		break
	// 	}
	// 	print(i)
	// }

	// fungsi continue akan men-skip perulangan saat itu dan akan loncat ke perulangan selanjutnya, fungsi dibawahnya tidak akan dijalankan
	// for i := 0; i < 10; i++ {
	// 	if i%2 == 0 {
	// 		continue
	// 	}
	// 	print(i)
	// }

	// ---------------- Function ------------------
	// function dibuat di luar scope function main, bisa didalamnya jika memang dibutuhkan.
	// func sayHello{
	// 	fmt.println("Hello")
	// }
	// sayHello()

	// ----------------- funciton parameter -----------------
	// jika ingin membuat function di dalam function main, masukkan function ke dalam variable.
	// paramter tidak jauh berbeda di javascript
	// sayHelloTo := func(firstName string, lastName string) {
	// 	fmt.Println("Hello", firstName, lastName)
	// }

	// sayHelloTo("Rafli", "Hamdani")

	//-------------------  Return Function -------------------
	// untuk mengembalikan nilai function, harus di tentukan terlebih dahulu tipe data return-nya. caranya dengan menuliskannya setelah tutup kurung parameter. ingat, kode program setelah return tidak akan dijalankan.
	// lebih baik nilai dari return dimasukkan kedalam variable
	// getHello := func(name string) string {
	// 	return "Hello " + name
	// }
	// fmt.Print(getHello("Rafli"))

	//------------------- Returning Multiple Values --------------

}
