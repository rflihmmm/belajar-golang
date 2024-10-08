package main

import (
	"errors"
	"fmt"
)

func main() {
	Pembagian := func(nilai int, pembagi int) (int, error) {
		if pembagi == 0 {
			return 0, errors.New("Pembagi dengan nol")
		} else {
			result := nilai / pembagi
			return result, nil
		}
	}

	hasil, err := Pembagian(7, 0)
	if err == nil {
		fmt.Println("Hasil", hasil)
	} else {
		fmt.Println("Error", err)
	}
}
