package main

import "fmt"

type Number interface {
	int64 | float64
}

func main() {
	// Initialize a map for the integer values
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	// Initialize a map for the float values
	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	fmt.Printf("Non-Generic Sums: %v and %v\n",
		SumInts(ints),
		SumFloats(floats),
	)

	fmt.Printf("Generic Sums: %v and %v\n",
		SumIntsOrFloats(ints),
		SumIntsOrFloats(floats),
	)

	fmt.Printf("Generic Sums with Constraint: %v and %v\n",
		SumNumbers(ints),
		SumNumbers(floats),
	)
}

// SumInts adds together the values of m.
func SumInts(ints map[string]int64) int64 {
	var sum int64
	for _, value := range ints {
		sum += value
	}
	return sum
}

// SumFloats adds together the values of m.
func SumFloats(ints map[string]float64) float64 {
	var sum float64
	for _, value := range ints {
		sum += value
	}
	return sum
}

// SumIntsOrFloats sums the values of map m. It supports both int64 and float64
// as types for map values.
func SumIntsOrFloats[KEY comparable, VALUE int64 | float64](ints map[KEY]VALUE) VALUE {
	var sum VALUE
	for _, value := range ints {
		sum += value
	}
	return sum
}

// SumNumbers sums the values of map m. It supports both integers
// and floats as map values.
func SumNumbers[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
