package main

import "fmt"

// Identity is a monad for a single value.
type Identity[T any] struct {
	Value T
}

// LiftingFunction transforms and lifts a single value: x -> y
type LiftingFunction[T any] func(T) Identity[T]

// Lift lifts a value into the Identity monad: x -> M(x)
func Lift[T any](value T) Identity[T] {
	return Identity[T]{value}
}

// FlatMap applies a LiftingFunction to the monad and unwraps it: M(x) -> y
func (i Identity[T]) FlatMap(f LiftingFunction[T]) T {
	return f(i.Value).Value
}

// Map applies a LiftingFunction to a monad: M(x) -> M(y)
func (i Identity[T]) Map(f LiftingFunction[T]) Identity[T] {
	return Lift(i.FlatMap(f))
}

// ComposeLifting composes two lifting functions into a single one.
// f(x): x -> M(y), g(x): x -> M(y) => f(g(x)): M(y)
func ComposeLifting[T any](f, g LiftingFunction[T]) LiftingFunction[T] {
	return func(x T) Identity[T] {
		return g(x).Map(f)
	}
}

func main() {
	// two lifting functions
	twiceLift := func(x int) Identity[int] {
		return Identity[int]{x * 2}
	}
	incLift := func(x int) Identity[int] {
		return Identity[int]{x + 1}
	}

	liftTwiceInc := ComposeLifting(twiceLift, incLift)
	fmt.Println(liftTwiceInc(20))
}
