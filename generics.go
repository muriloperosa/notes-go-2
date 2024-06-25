package main

func add[T string | int | float64](a, b T) T {
	return a + b
}
