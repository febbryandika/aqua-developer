package main

import (
	"fmt"
	"sort"
)

type ListProduk struct {
	nama  string
	harga int32
}

type ListProduct []ListProduk

// Jawaban poin 1.b
func (list ListProduct) produkTermurah() (string, int32) {
	var count int32 = 100000
	var name string
	for _, produk := range list {
		if produk.harga < count {
			count = produk.harga
			name = produk.nama
		} else {
			continue
		}
	}
	return name, count
}

// Jawaban poin 1.b
func (list ListProduct) produkTermahal() (string, int32) {
	var count int32
	var name string
	for _, produk := range list {
		if produk.harga > count {
			count = produk.harga
			name = produk.nama
		} else {
			continue
		}
	}
	return name, count
}

// Jawaban poin 1.c
func (list ListProduct) sepuluhRibu() {
	for _, produk := range list {
		if produk.harga == 10000 {
			fmt.Printf("%s - %d\n", produk.nama, produk.harga)
		} else {
			continue
		}
	}
}

func separator() {
	fmt.Println("----------------------------------------------------------------")
}

func main() {
	var list = []ListProduk{
		{"Benih Lele", 50000},
		{"Pakan Lele Cap Menara", 25000},
		{"Probiotik A", 75000},
		{"Probiotik Nila B", 10000},
		{"Pakan Nila", 20000},
		{"Benih Nila", 20000},
		{"Cupang", 5000},
		{"Benih Nila", 30000},
		{"Benih Cupang", 10000},
		{"Probiotik B", 10000},
	}

	var uang int32 = 100000
	var total int32
	var index = 0

	// Jawaban poin 1.a
	sortedList := []ListProduk{}
	for i := 0; i < len(list); i++ {
		sortedList = append(sortedList, list[i])
	}
	sort.Slice(sortedList, func(i, j int) bool {
		return sortedList[i].harga < sortedList[j].harga
	})
	for i, produk := range sortedList {
		uang -= produk.harga
		if uang >= 0 {
			total += produk.harga
		} else {
			index = i
			break
		}
	}
	fmt.Println("Total produk dengan harga dibawah  Rp 100.000:")
	fmt.Printf("Harga total: %d\n", total)
	fmt.Printf("Jumlah barang yang dibeli: %d\n", index)
	fmt.Println("Produk yang dibeli: ")
	for i := 0; i < index; i++ {
		fmt.Printf("%s - %d\n", sortedList[i].nama, sortedList[i].harga)
	}

	separator()

	// Jawaban poin 1.c
	fmt.Println("Daftar produk dengan harga Rp 10.000 :")
	ListProduct.sepuluhRibu(list)

	separator()

	// Jawaban poin 1.b
	namaProdukTermurah, hargaProdukTermurah := ListProduct.produkTermurah(list)
	fmt.Printf("Daftar produk termurah: %s Rp %d\n", namaProdukTermurah, hargaProdukTermurah)
	namaProdukTermahal, hargaProdukTermahal := ListProduct.produkTermahal(list)
	fmt.Printf("Daftar produk termahal: %s Rp %d\n", namaProdukTermahal, hargaProdukTermahal)
}
