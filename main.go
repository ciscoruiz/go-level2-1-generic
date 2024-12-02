package main

import (
	"fmt"

	"github.com/ciscoruiz/go-level2-1-generic/generic"
)

func equalToFiveInt(value int) bool {
	return value == 5
}
func equalToTenInt(value int) bool {
	return value == 10
}
func equalToTwentyTwoInt(value int) bool {
	return value == 22
}
func equalToOneHundredInt(value int) bool {
	return value == 100
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

	fmt.Printf("Exists 5: %v\n", generic.AnyOf(integers, equalToFiveInt))
	fmt.Printf("Exists 10: %v\n", generic.AnyOf(integers, equalToTenInt))
	fmt.Printf("Greather that 5: %v\n", generic.AnyOf(integers, greatherThanFiveInt))

	fmt.Printf("Index for 5: %v\n", generic.FindIf(integers, equalToFiveInt))
	fmt.Printf("Index for 10: %v\n", generic.FindIf(integers, equalToTenInt))

	adjacents := []int{10, 11, 12, 13, 14, 15}
	fmt.Printf("Index-adjacent for adjacents: %v\n", generic.AdjacentFind(adjacents, areAdjacentsPlusOne))
	fmt.Printf("Index-adjacent for adjacents-ten: %v\n", generic.AdjacentFind(integers, areAdjacentsPlusTen))
	fmt.Printf("Index-adjacent for adjacents: %v\n", generic.AdjacentFind(adjacents, areAdjacentsPlusTen))

	fmt.Printf("Are equals: %v\n", generic.Equal(adjacents, integers))
	fmt.Printf("Are equals: %v\n", generic.Equal(adjacents, adjacents))

	fmt.Printf("Replace by 100: %v\n", generic.ReplaceIf(adjacents, 22, equalToOneHundredInt))
	fmt.Printf("Replace by 10: %v values=%v\n", generic.ReplaceIf(integers, 22, equalToTenInt), integers)

	fmt.Printf("Remove-if 100: %v\n", generic.RemoveIf(adjacents, equalToOneHundredInt))
	fmt.Printf("Remove by 22: %v values=%v\n", generic.RemoveIf(integers, equalToTwentyTwoInt), integers)

	fmt.Printf("Sorted: %v result=%v\n", adjacents, generic.IsSorted(adjacents))
	fmt.Printf("Sorted: %v result=%v\n", integers, generic.IsSorted(integers))
}
