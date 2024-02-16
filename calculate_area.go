package main

import (
	"fmt"
	"math"
	"strings"
)

/*
	Hint: Use the following formulas to calculate the area of different shapes.

	Rectangle: length * width
	Circle: π * radius * radius
	Triangle: 0.5 * base * height
	Square: side * side

	Note: You can use math.Pi for π and math.Pow(x, y) for x^y.
*/

// Area calculates the area of different shapes based on the shape and provided dimensions.
func Area(shape string, dimensions []float64) float64 {
	// TODO 1: Create a variable named area of type float64 and assign 0 to it.
	var area float64 = 0

	// TODO 2: Use a switch statement or if statement to calculate the area of the shape.
	switch strings.ToLower(shape) {
	case "rectangle":
		if len(dimensions) != 2 {
			return -1.0 // Validasi Persegi panjang
		}
		length := dimensions[0]
		width := dimensions[1]
		area = length * width
	case "circle":
		if len(dimensions) != 1 {
			return -1.0 // Validasi Lingkaran
		}
		radius := dimensions[0]
		area = math.Pi * math.Pow(radius, 2)
	case "triangle":
		if len(dimensions) != 2 {
			return -1.0 // Validasi Triangle
		}
		base := dimensions[0]
		height := dimensions[1]
		area = 0.5 * base * height
	case "square":
		if len(dimensions) != 1 {
			return -1.0 // Validasi Square
		}
		side := dimensions[0]
		area = math.Pow(side, 2)
	default:
		return -1.0 // Tidak Ada Bangun Datar
	}

	// TODO 3: Return the area variable.
	return area
}

// Menambahkan Fungsi Main Untuk Menjalankan Fungsi Area dari Program File calculate_area.go
func main() {
	// Menghitung dan mencetak luas persegi panjang.
	rectangleArea := Area("Rectangle", []float64{5, 3})
	fmt.Printf("Luas persegi panjang: %.2f\n", rectangleArea)

	// Menghitung dan mencetak luas lingkaran.
	circleArea := Area("Circle", []float64{4})
	fmt.Printf("Luas lingkaran: %.2f\n", circleArea)

	// Menghitung dan mencetak luas segitiga.
	triangleArea := Area("Triangle", []float64{6, 8})
	fmt.Printf("Luas segitiga: %.2f\n", triangleArea)

	// Menghitung dan mencetak luas persegi.
	squareArea := Area("Square", []float64{5})
	fmt.Printf("Luas persegi: %.2f\n", squareArea)
}
