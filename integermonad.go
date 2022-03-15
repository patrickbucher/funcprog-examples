package main

import "fmt"

type Int int
type Op func(Int, Int) Int
type OpErr func(Int, Int) (Int, error)

func Unit(x Int) (Int, error) {
	return x, nil
}

func Lift(f func(Int, Int) Int) func(Int, Int) (Int, error) {
	return func(x Int, y Int) (Int, error) {
		return Unit(f(x, y))
	}
}

func Bind(

func Mul(x, y Int) Int {
	return x * y
}

func Div(x, y Int) (Int, error) {
	if y == 0 {
		return 0, fmt.Errorf("%d / %d: divide by zero", y, x)
	}
	return x / y
}

func main() {
	fmt.Println(Div(Mul(3, 2), 2))
	fmt.Println(Mul(Div(6, 2), 2))
	fmt.Println(Mul(Div(6, 0), 2))
}
