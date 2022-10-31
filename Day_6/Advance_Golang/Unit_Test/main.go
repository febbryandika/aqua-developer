package main

import (
	"fmt"
	"math"

	log "github.com/sirupsen/logrus"
)

type Square struct {
	Side float64
}

func (s Square) Area() float64 {
	return math.Pow(s.Side, 2)
}

func (s Square) Circumference() float64 {
	return s.Side * 4
}

func main() {
	twoDimensional := Square{10}
	log.SetFormatter(&log.JSONFormatter{})
	log.Printf("Run Square Calculator")
	fmt.Printf("Square Area: %.f\n", twoDimensional.Area())
	fmt.Printf("Square Circumference: %.f\n", twoDimensional.Circumference())
	fmt.Println()
}
