package main

import (
	"errors"
	"fmt"
	"strings"
)

type FloatOp func(float64) float64
type FloatErrOp func(float64) (float64, error)

func Square(x float64) float64 {
	return x * x
}

func Cube(x float64) float64 {
	return x * x * x
}

func Reciprocal(x float64) (float64, error) {
	if x == 0 {
		return 0.0, errors.New("divide by zero")
	}
	return 1 / x, nil
}

func Compose(f, g FloatOp) FloatOp {
	return func(x float64) float64 {
		return f(g(x))
	}
}

func ComposeErr(f, g FloatErrOp) FloatErrOp {
	return func(x float64) (float64, error) {
		gx, errG := g(x)
		fx, errF := f(gx)
		return fx, ToErr(errG, errF)
	}
}

func ComposeManyErr(fs ...FloatErrOp) FloatErrOp {
	return func(x float64) (float64, error) {
		errs := make([]error, 0)
		for _, f := range fs {
			fx, err := f(x)
			errs = append(errs, err)
			x = fx
		}
		return x, ToErr(errs...)
	}
}

func ToErr(errs ...error) error {
	strs := make([]string, 0)
	for _, err := range errs {
		if err != nil {
			strs = append(strs, fmt.Sprintf("%v", err))
		}
	}
	return errors.New(strings.Join(strs, ": "))
}

func Lift(f FloatOp) FloatErrOp {
	return func(x float64) (float64, error) {
		return f(x), nil
	}
}

func main() {
	fmt.Println(Square(2))
	fmt.Println(Cube(2))

	squaredCube := Compose(Square, Cube)
	fmt.Println(squaredCube(2))

	fmt.Println(Reciprocal(2))
	fmt.Println(Reciprocal(1))
	fmt.Println(Reciprocal(0))

	squaredReciprocal := ComposeErr(Lift(Square), Reciprocal)
	fmt.Println(squaredReciprocal(5))

	reciprocalAndBack := ComposeManyErr(Reciprocal, Reciprocal)
	fmt.Println(reciprocalAndBack(3))
	fmt.Println(reciprocalAndBack(0))
}
