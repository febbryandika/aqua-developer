package main

import (
	"fmt"
)

// Functions
func swap(x, y string) (string, string) {
	return y, x
}

// Struct

type person struct {
	name string
	age  int
}

type student struct {
	mataKuliah string
	person
}

func main() {
	//
	fmt.Println("Hello world!")
	fmt.Println()

	// Number
	nama := "Febbry"
	umur := 20
	decimalNumber := 3.14
	fmt.Printf("Nama saya %s, umur saya %d\n", nama, umur)
	fmt.Printf("1 angka dibelakang koma %.1f\n", decimalNumber)
	fmt.Printf("1 angka dibelakang koma %.2f\n", decimalNumber)
	fmt.Printf("1 angka dibelakang koma %.3f\n", decimalNumber)
	fmt.Println()

	// Boolean
	exist := false
	fmt.Printf("Apakah ada? %t\n\n", exist)
	fmt.Println()

	// Multiple line strings
	message := `Hello!
	My Name is
	Andika`
	fmt.Println(message)
	fmt.Println()

	// Variable declarations
	word := "Hello"
	word = "Hi"
	fmt.Println(word)
	fmt.Println()

	// Declarations multiple variables
	firstName, lastName := "Febbry", "Andika Ramadhan"
	fmt.Println(firstName, lastName)
	fmt.Println()

	// Constants
	const tahun int = 2022
	fmt.Println(tahun)

	// Functions
	x, y := swap("Febbry", "Andika")
	fmt.Println(x, y)
	fmt.Println()

	// Struct
	var muridPertama student
	muridPertama.name = "Febbry"
	muridPertama.age = 20
	muridPertama.mataKuliah = "Golang"
	fmt.Println(muridPertama)
	fmt.Println("Nama : ", muridPertama.name)
	fmt.Println("Umur : ", muridPertama.age)
	fmt.Println("Mata Kuliah : ", muridPertama.mataKuliah)
	fmt.Println()

	var muridKedua = student{
		mataKuliah: "JavaScript",
		person: person{
			name: "Andika",
			age:  20,
		},
	}
	fmt.Println(muridKedua)
	fmt.Println("Nama : ", muridKedua.name)
	fmt.Println("Umur : ", muridKedua.age)
	fmt.Println("Mata Kuliah : ", muridKedua.mataKuliah)
	fmt.Println()

	var murid = []student{
		{
			mataKuliah: "C",
			person: person{
				name: "A",
				age:  20,
			},
		},
		{
			mataKuliah: "Ruby",
			person: person{
				name: "B",
				age:  20,
			},
		},
		{
			mataKuliah: "Java",
			person: person{
				name: "C",
				age:  20,
			},
		},
	}
	for _, student := range murid {
		fmt.Println(student.name, student.age)
	}
	fmt.Println()

	// Data Structures
	ayam := map[string]int{}
	ayam["A"] = 1
	ayam["B"] = 2
	fmt.Println("Ayam A : ", ayam["A"])
	fmt.Println("Ayam B : ", ayam["B"])
	fmt.Println()

	var namaSendiri [3]string
	namaSendiri[0] = "Febbry"
	namaSendiri[1] = "Andika"
	namaSendiri[2] = "Ramadhan"
	fmt.Println(namaSendiri)
	fmt.Println(namaSendiri[0], namaSendiri[1], namaSendiri[2])
	fmt.Println()

	// Slices
	buah := []string{"Apel", "Nanas", "Mangga", "Jeruk", "Sirsak"}
	fmt.Println(buah)
	fmt.Println(buah[0:2])
	fmt.Println(buah[2:])
	fmt.Println(buah[:2])
	fmt.Println(buah[:])
	fmt.Println(buah[0:0])
	fmt.Println(buah[2:2])
	fmt.Println(buah[1:4])
	fmt.Println()

	// Defer
	defer fmt.Println("Belajar Basic Golang")
	fmt.Println("Hello, world!")
}
