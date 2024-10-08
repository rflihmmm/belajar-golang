// --------------------- Interface --------------------
/**
  Interface adalah tipe data Abstract, dia tidak memiliki impelementasi langsung
  sebuah interface berisikan definisi-definisi method
  Biasanya interface digunakan sebagai kontrak
*/

// Implementasi Interface
/**
  Setiap tipe data yang sesuai dengan kontrak interface, secara otomatis dianggap sebagai interface tersebut
sehingga kita tidak perlu mengimplementasikan interface secara manual
Hal ini agar berbeda dengan bahasa pemrograman lain yang ketika membuat interface, kita harus menyebutkan secara eksplisit akan menggunakan interface yang mana
*/
package main

import "fmt"

type Person struct {
	Name string
}

type HasName interface {
	GetName() string
}

func sayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

func (person Person) GetName() string {
	return person.Name
}

func main() {
	var eko Person
	eko.Name = "Eko"
	sayHello(eko)
}
