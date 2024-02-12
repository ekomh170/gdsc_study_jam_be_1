package main

import "fmt"

func main() {
	// menampilkan hallo dunia di console
	fmt.Println("Hello, Dunia!")

	// Variabel Untuk Menyimpan Data
	// var name string
	// var age int

	// name = "Eko Muchamad Haryono"
	// age = 20

	// fmt.Println(name)
	// fmt.Println(age)
	// var nim, alamat string

	// Variabel Tidak Dapat Di Rubah Menggunakan Fungsi Const
	// const nim = "123456789"
	// nim = "1234567890"

	// TYPE Deklarasi Mirip Kayak AS di SQL gak Sih?
	type KtmNF string
	var name7 KtmNF = "Eko Muchamad Haryono"
	fmt.Println(name7)
	fmt.Println(KtmNF("1234567890"))

	// Math Operasion
	// var a = 10
	// var b = 10
	// var c = a + b
	// fmt.Println(c)

	// Augmented Assignments
	// Basic Math Operation
	// var a = 10
	// a += 10
	//output a = 20
	// fmt.Println(a)
	// var b = 10
	// b -= 10
	//output b = 0
	// fmt.Println(b)
	//etc

	// Array
	var names [3]string
	names[0] = "Eko"
	names[1] = "Muchamad"
	names[2] = "Haryono"

	fmt.Println(names[0], names[1], names[2])

	// Slices
	myslice := names[1:2]
	fmt.Println("My Name is = ", myslice)

	// Maps
	person := map[string]string{"name": "Eko", "address": "Bogor"}
	fmt.Println(person)
	fmt.Println("My Name is", person["name"])
	fmt.Println("My Address is", person["address"])

	// If Expression
	name := "Eko"
	if name == "Haryono" {
		fmt.Println("Name : Eko")
	} else {
		fmt.Println("Siapa Nama Kamu")
	}

	// Switch Expression
	switch name {
	case "Eko":
		fmt.Println("Name : Eko")
	case "Haryono":
		fmt.Println("Name : Haryono")
	default:
		fmt.Println("Nama Tidak Di Temukan")
	}

	// Looping

	// For Range
	data := []string{"Eko", "Muhammad", "Haryono"}
	for index, name := range data {
		fmt.Println("Index ", index, "= ", name)
	}

	// Break

	// Continue

	// Manggil Fungsi di Fungsi Utama
	Hello("Eko\n\n")

	struct_test()

}

// Function

func Hello(name string) {
	fmt.Println("hai", name)
}

// Pointer

func pointer() {
	var number int = 5               //deklarasi variabel number dengan nilai 5
	var numberPointer *int = &number //deklarasi variabel numberPointer sebagai pointer dari variabel number

	fmt.Println("Nilai dari number:", number)
	fmt.Println("Alamat memory dari number:", &number)
	fmt.Println("Nilai dari numberPointer:", *numberPointer)
	fmt.Println("Alamat memory dari numberPointer:", numberPointer)
}

// Struct

type Person struct {
	name    string
	age     int
	address string
}

func struct_test() {
	person := Person{name: "Asnur", age: 21, address: "Bogor"}
	fmt.Println("Nama:", person.name)
	fmt.Println("Usia:", person.age)
	fmt.Println("Alamat:", person.address)
}

// Struct Method

func (p Person) setAddress(newAddress string) {
	p.address = newAddress
}

func (p Person) getAddress() string {
	return p.address
}
