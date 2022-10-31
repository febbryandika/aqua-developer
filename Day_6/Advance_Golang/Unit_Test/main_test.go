package main

import "testing"

var (
	side                Square  = Square{10}
	squareArea          float64 = 100
	squareCircumference float64 = 40
)

func TestSquareArea(t *testing.T) {
	t.Logf("%.2f", side.Area())

	if side.Area() != squareArea {
		t.Errorf("SALAH! harusnya %.2f", squareArea)
	}

}

func TestCircleCircumference(t *testing.T) {
	t.Logf("%.2f", side.Circumference())

	if side.Circumference() != squareCircumference {
		t.Errorf("SALAH! harusnya %.2f", squareCircumference)
	}

}
