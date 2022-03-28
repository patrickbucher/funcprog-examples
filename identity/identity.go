package main

import "fmt"

// Identity is a functor for a single value.
type Identity[T any] struct {
	Value T
}

// Function transforms a single value.
type Function[T any] func(T) T

// Map performs the given Function on the Identity.
func (i Identity[T]) Map(f Function[T]) Identity[T] {
	return Identity[T]{Value: f(i.Value)}
}

// Compose composes the functions f and g, returning a function f(g(x)).
func Compose[T any](f, g Function[T]) Function[T] {
	return func(x T) T { return f(g(x)) }
}

func main() {
	// 1st Functor Law: Identity
	identity := func(v int) int { return v }

	a := Identity[int]{3}
	b := a.Map(identity)
	fmt.Println(a, b)

	// 2nd Functor Law: Composition
	twice := func(v int) int { return v * 2 }
	inc := func(v int) int { return v + 1 }

	m := Identity[int]{3}
	n := m.Map(twice).Map(inc)
	fmt.Println(m, n)

	x := Identity[int]{3}
	y := x.Map(Compose(inc, twice))
	fmt.Println(x, y)
}
