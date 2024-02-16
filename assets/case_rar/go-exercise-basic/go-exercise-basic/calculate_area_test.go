package main

import (
	"math"
	"testing"
)

func TestCalcShapeArea(t *testing.T) {
	var testCases = []struct {
		shape       string
		dimensions  []float64
		expected    float64
		description string
	}{
		{"rectangle", []float64{5, 10}, 50.0, "Valid Rectangle"},
		{"circle", []float64{7}, float64(math.Pi * math.Pow(7, 2)), "Valid Circle"},
		{"triangle", []float64{8, 4}, 16.0, "Valid Triangle"},
		{"square", []float64{5}, 25.0, "Valid Square"},
		{"rectangle", []float64{5}, -1.0, "Invalid Rectangle"},
		{"circle", []float64{7, 3}, -1.0, "Invalid Circle"},
		{"triangle", []float64{8}, -1.0, "Invalid Triangle"},
		{"unknown", []float64{5, 10}, -1.0, "Unknown Shape"},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			actual := Area(testCase.shape, testCase.dimensions)
			if actual != testCase.expected {
				t.Errorf("Expected %f, got %f", testCase.expected, actual)
			}
		})
	}
}
