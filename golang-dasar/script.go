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

	// fmt.Println(name)
	// fmt.Println(e)
	// fmt.Println(bulat)

	// var name [2] string
	// name[0] = "Rafli"
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
	// potongan data array. slice mengacu pada array, jadi jika array diubah, maka slice pun akan berubah. Begitupun sebaliknya, jika slice diubah, array pada slice tersebut akan berubah pula.

	// bulan := [...]string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}

	// slice1 akan mengambil array bulan index ke 4 sampai sebelum index ke 7. jika bulan[4:], berarti dari index ke 4 samapi index terakhir. dan jika bulan[:7], berarti dari index paling awal samapai sebelum index ke 7.
	// slice1 := bulan[4:7]
	// fmt.Println(slice1) // ["May", "June", "July"]
	// slice1[1] = "Mei" // ["May", "Mei", "July"], array bulan juga akan beruba
	// fmt.Println(len(slice1)) // 3
	// fmt.Println(cap(slice1)) // 8 ---- cap berarti kapasitas slice tersebut. maksudnya panjang slice yang bisa didapatkan yang mengacu pada array.

	// slice2 := bulan[10:]
	// slice3 := append(slice2, "duarsember") // membuat slice baru yang datanya diambil dari slice 2, parameter kedua adalah data yang ingin ditambahkan di index paling akhir.
	// fmt.Println(slice3) // slice3 tidak akan merubah array utama dan akan membuat array baru. itu dikarenakan slice3 melebihi kapasitas dari slice2.

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
	// map[tipedatakey]tipedatavalue{}
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

	// If Short Statement
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

	// fungsi continue akan men-skip perulangan saat itu dan akan loncat ke perulangan selanjutnya, fungsi dibawahnya tidak akan dijalankan atau diabaikan dan di lanjutkan ke perulangan selanjutnya
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
	// // mengembalikan atau me-return value lebih dari satu
	// getName := func() (string, int, string) {
	// 	nama := "Muh Rafli Hamdani"
	// 	umur := 23
	// 	alamat := "Makassar"
	//
	// 	return nama, umur, alamat
	// }
	//
	// nama, umur, alamat := getName()
	//
	// fmt.Println(
	// 	"Halo, nama saya "+
	// 		nama+
	// 		", umur saya ",
	// 	umur,
	// 	", saya tinggal di kota",
	// 	alamat+
	// 		".",
	// )

	//-------------------- Named Return Values ---------------------
	// // memberi nama retrun vale sebagai varible yang dimana variable tersebut dideklarasikan di dalam function tersebut
	// getCompleteName := func() (firstName, midName, lastName string, umur int) {
	// 	firstName = "Muh"
	// 	midName = "Rafli"
	// 	lastName = "Hamdani"
	// 	umur = 21
	//
	// return firstName, midName, lastName, umur
	// bisa melakukan retrun secara explisit seperti diatas atau bisa juga dengan cukup mengetikan retrun
	// 	return
	// }
	//
	// namaDepan, namaTengah, namaBelakang, umur := getCompleteName()
	//
	// fmt.Println(namaDepan, namaTengah, namaBelakang, "umur", umur)

	// ----------------- Variadic Function ----------------------
	// parameter atau argumen yang terletak di akhir pada golang biasanya di sebut varags (variable argumen)
	// dalam datu argumen tersebut, kita dapat menerima lebih dari satu input atau bisa juga dianggap seperti array
	// bedanya parameter biasa dengan tipe data array adalah tipe data array kita wajib membuat array sebelum mengerimnya ke function sedangkan parameter yang menggunakan varags dapat langsung mengirim datanya, jika lebih dari satu, cukup gunakan tanda koma

	// getTotal := func(numbers ...int) int {
	// 	total := 0
	// 	for _, number := range numbers {
	// 		total += number
	// 	}
	// 	return total
	// }
	//
	// total := getTotal(10, 3, 14, 15, 12)
	// fmt.Println(total)

	// ------ Slice parameter -----
	// numbers := []int{10, 12, 4, 12}
	// total := getTotal(numbers...)
	// fmt.Println(total)

	//---------------------- Function Value ---------------
	// function dapat di masukkan ke dalam variable
	// getGoodBye := func(name string) string {
	// 	return "Hello " + name
	// }
	//
	// goodBye := getGoodBye
	// fmt.Println(goodBye("Rafli"))

	// ------------------- Function as Parameter -------------
	// function juga dapat di jadikan parameter pada function lain
	// spamFilter := func(name string) string {
	// 	if name == "Anjing" {
	// 		return "..."
	// 	} else {
	// 		return name
	// 	}
	// }
	//
	// sayButFilter := func(name string, filter func(string) string) {
	// 	fmt.Println("Hello", filter(name))
	// }
	//
	// sayButFilter("Anjing", spamFilter)
	// sayButFilter("Rafli", spamFilter)\

	//-------------------- Anonymous Fucntion ---------------
	// intinya Anonymous Function adalah function tanpa nama. seperti di dideklarasikan sebagai variable atau dibuat langsung sebagai parameter pada fucntion lainnya

	// type BlackList func(string) bool
	// function didalam variable
	// registerUser := func(name string, blacklist BlackList) {
	// 	if blacklist(name) {
	// 		fmt.Println("You are blocked", name)
	// 	} else {
	// 		fmt.Println("Welcome", name)
	// 	}
	// }
	//
	// blacklist := func(name string) bool {
	// 	return name == "Anjing"
	// }
	//
	// registerUser("Rafli", blacklist)
	//
	// function yang dibuat langsung sebagai parameter
	// registerUser("Anjing", func(name string) bool {
	// 	return name == "Anjing"
	// })

	// ------------------ Recursive Function ------------
	// function yang memanggil dirinya sendiri atau looping
	// contoh kasus adalah factorial
	// menggunakan For Loops

	// factorialLoop := func(value int) int {
	// 	result := 1
	// 	for i := value; i > 0; i-- {
	// 		result *= i
	// 	}
	// 	return result
	// }
	//
	// fmt.Println(factorialLoop(5))
	//
	// Recursive function tidak dapat digunakan jika function tersebut adalah Anonymous
	// fmt.Println(factorialRecursive(5))

	// --------------------- Closure -----------------
	// kemampuan sebuah function berinteraksi dengan data-data di sekitarnya dalam scope yang sama
	// hati-hati saat menggunaka closure ini

	// counter := 0
	// name := "budi"
	// increment := func() {
	// variable dapat di ubah didalam scope ini, tapi tidak jika mendeklarisikan ulang variable tersebut walaupun dengan nama yang sama
	// name := "Rafli"
	// fmt.Println("increment")
	// counter++
	// varible counter yang berada diluar function increment() dapat diakses
	// }

	// varible yang berada di dalam function increment tidak dapat di akses
	//
	// increment()
	// increment()
	// fmt.Println(counter)
	// fmt.Println(name)

	// ----------------- Defer, Panic dan Recover -------------
	// ---------- Defer Function
	// defer function adalah function yang bisa dijadwalkan untuk di eksekusi setelah sebuah function selesai di eksekusi
	// defer function akan selalu di eksekusi walaupun terjadi error di function yang dieksekusi
	// logging := func() {
	// fmt.Println("Selesai memanggil function")
	// }
	// defer function selalu di tempatkan paling atas setelah fungsi dibuat, jika di tempatkan di paling bawah dan terjadi error, defer function tidak dapat dijalankan atau di eksekusi
	// runApp := func(value int) {
	// defer logging()
	// fmt.Println("Run Application")
	// result := 10 / value
	// fmt.Println("Reseult", result)
	// }

	// runApp(0)

	//--------------- Panic
	// Panic function digunakan jika ingin menghentikan program yang sedang berjalan
	// Panic function biasanya dipanggil ketika terjadi error pada saat program berjalan
	// Saat Panic function dipanggil, program akan terhenti, namun defer function tetap akan di eksekusi
	// endApp := func() {
	// 	fmt.Println("End App")
	// }
	//
	// runApp := func(error bool) {
	// 	defer endApp()
	// 	if error {
	// 		panic("Error")
	// 	}
	// 	fmt.Println("App is Running")
	// }
	//
	// runApp(true)

	//---------------- Recover
	// Recover adalah function yang bisa kita gunakan untuk menangkap data panic
	// Dengan recover proses panic akan terhenti, sehingga program akan tetap berjalan

	/**
	  endApp := func() {
			message := recover()
			if message != nil {
				fmt.Println("Error dengan message", message)
			}
			fmt.Println("End App")
		}

		runApp := func(error bool) {
			defer endApp()
			if error {
				panic("Error")
			}
			fmt.Println("App is Running")
		}

		runApp(true)
		fmt.Println("rafli")
	*/

	//------------------------ Struct ------------------------
	/** Struct adalah sebuah template data yang digunakan untuk menggabungkan nol atau lebih tipe data lainnya dalam satu kesatuan
	struct biasanya representasikan data dalam program aplikasi yang kita buat
	Data di struct disimpan dalam field
	sederhananya struct adalah kumpulan dari field
	*/
	/**
		  type Customer struct {
				Name, Address string
				Age           int
	}
			// struct tidak bisa langsung digunakan
			// Namun kita bisa membuat data/object dari struct yang telah dibuat
			var rafli Customer
			rafli.Name = "Rafli Hamdani"
			rafli.Address = "Makassar"
			rafli.Age = 20

			fmt.Println(rafli) // {Rafli Hamdani Makassar 20}

			//----------------- Struct Literals -------------
			joko := Customer{
				Name:    "Joko",
				Address: "Cirebon",
				Age:     20,
			}
			fmt.Println(joko)
	*/

	//---------------------- Struct Method --------------
	/**
	  Struct adalah tipe data seperti tipe data lainnya, dia bisa digunakan sebagai parameter untuk function
	  namun jika ingin menambahkan method ke dalam structs, shinggak seakan-akan sebuah struct memiliki function
	  Method adalah function
	*/
	/**
	  func (customer Customer) sayHi(name string) {
		  fmt.Println("Hello", name, "My nase is", customer.Name)
	  }

		Rafli.sayHi("Joko") */
}

// ------------------ recursive function ------------
// func factorialRecursive(value int) int {
// 	if value == 1 {
// 		return 1
// 	} else {
// 		return value * factorialRecursive(value-1)
// 	}
// }
