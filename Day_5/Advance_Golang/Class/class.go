package main

import (
	"fmt"
	"math"
	"runtime"
)

// Goroutines
func print(till int, message string) {
	for i := 1; i <= till; i++ {
		fmt.Println(i, message)
	}
}

// Interface
type calculate interface {
	area() float64
	circumference() float64
}

type circle struct {
	radius float64
}

type square struct {
	side float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c circle) circumference() float64 {
	return math.Pi * c.radius * 2
}

func (s square) area() float64 {
	return math.Pow(s.side, 2)
}

func (s square) circumference() float64 {
	return s.side * 4
}

// Method
type student struct {
	name string
	age  int
}

func (s student) sayHello() {
	fmt.Printf("Hello, %s! You are %d years old.\n", s.name, s.age)
}

func (s *student) changeName(name string) {
	s.name = name
}

func main() {
	// Pointer
	var x int = 34289
	var p *int = &x
	fmt.Println(*p)
	fmt.Println("Value stored in x : ", x)
	fmt.Println("Address of x : ", &x)
	fmt.Println("Value stored in p : ", p)
	fmt.Println()

	var numberA int = 4
	var numberB *int = &numberA
	fmt.Println("Value stored in numberA : ", numberA)
	fmt.Println("Address of numberA : ", &numberA)
	fmt.Println("Value stored in numberB : ", *numberB)
	fmt.Println("Address of numberB : ", &numberB)
	fmt.Println()
	numberA = 10
	fmt.Println("Value stored in numberA : ", numberA)
	fmt.Println("Address of numberA : ", &numberA)
	fmt.Println("Value stored in numberB : ", *numberB)
	fmt.Println("Address of numberB : ", &numberB)
	fmt.Println()

	// Method
	var student1 = student{"Febbry", 10}
	student1.sayHello()
	student1.changeName("Andika")
	student1.sayHello()

	// Interface
	var twoDimensional calculate

	twoDimensional = square{10}
	fmt.Println(twoDimensional.area())
	fmt.Println(twoDimensional.circumference())
	fmt.Println()

	twoDimensional = circle{14}
	fmt.Println(twoDimensional.area())
	fmt.Println(twoDimensional.circumference())
	fmt.Println()

	//Goroutines
	runtime.GOMAXPROCS(2)

	go print(5, "Hello")
	print(5, "Hi")

	var input string
	fmt.Scanln(&input)
}
