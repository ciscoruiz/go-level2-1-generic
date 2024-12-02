package main

import "fmt"

func AnyOf[T comparable](values []T, yield func(T) bool) bool {
	for _, t := range values {
		if yield(t) {
			return true
		}
	}
	return false
}

func FindIf[T comparable](values []T, yield func(T) bool) int {
	for index, t := range values {
		if yield(t) {
			return index
		}
	}
	return -1
}

func AdjacentFind[T comparable](values []T, yield func(T, T) bool) int {
	len := len(values)

	if len == 0 {
		return -1
	}

	previous := values[0]

	for index := 1; index < len; index++ {
		if yield(previous, values[index]) {
			return index
		}
		previous = values[index]
	}
	return -1
}

func equalToFiveInt(value int) bool {
	return value == 5
}
func equalToTenInt(value int) bool {
	return value == 10
}

func greatherThanFiveInt(value int) bool {
	return value > 5
}

func areAdjacentsPlusOne(first int, second int) bool {
	return (first + 1) == second
}
func areAdjacentsPlusTen(first int, second int) bool {
	return (first + 10) == second
}

func main() {
	integers := []int{10, 20, 30, 50, 10, 20}

	fmt.Printf("Exists 5: %v\n", AnyOf(integers, equalToFiveInt))
	fmt.Printf("Exists 10: %v\n", AnyOf(integers, equalToTenInt))
	fmt.Printf("Greather that 5: %v\n", AnyOf(integers, greatherThanFiveInt))

	fmt.Printf("Index for 5: %v\n", FindIf(integers, equalToFiveInt))
	fmt.Printf("Index for 10: %v\n", FindIf(integers, equalToTenInt))

	adjacents := []int{10, 11, 12, 13, 14, 15}
	fmt.Printf("Index-adjacent for adjacents: %v\n", AdjacentFind(adjacents, areAdjacentsPlusOne))
	fmt.Printf("Index-adjacent for adjacents-ten: %v\n", AdjacentFind(integers, areAdjacentsPlusTen))
	fmt.Printf("Index-adjacent for adjacents: %v\n", AdjacentFind(adjacents, areAdjacentsPlusTen))

}
