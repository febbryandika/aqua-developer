package main

import "fmt"

func rectangleStar(number int) string {
	var star string
	for i := 0; i < number; i++ {
		for j := 0; j < number; j++ {
			star += "*"
		}
		star += "\n"
	}
	return star
}

func triangleStar(number int) string {
	var star string
	for i := 0; i <= number; i++ {
		for j := 0; j < i; j++ {
			star += "*"
		}
		star += "\n"
	}
	return star
}

func main() {
	var number int
	fmt.Print("Masukan angka : ")
	fmt.Scan(&number)
	fmt.Println("================================================")
	if number%2 == 0 {
		fmt.Println("Ini bentuk jika angka genap")
		fmt.Println(rectangleStar(number))
	} else {
		fmt.Println("Ini bentuk jika angka ganjil")
		fmt.Println(triangleStar(number))
	}
}
