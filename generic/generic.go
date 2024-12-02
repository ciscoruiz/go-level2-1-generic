package generic

import "golang.org/x/exp/constraints"

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

func Equal[T comparable](first []T, second []T) bool {
	if len(first) != len(second) {
		return false
	}

	for index := 0; index < len(first); index++ {
		if first[index] != second[index] {
			return false
		}
	}

	return true
}

func ReplaceIf[T comparable](values []T, new_value T, yield func(T) bool) int {
	counter := 0

	for index := 0; index < len(values); index++ {
		if yield(values[index]) {
			values[index] = new_value
			counter++
		}

	}

	return counter
}

func RemoveIf[T comparable](values []T, yield func(T) bool) []T {
	counter := 0

	result := values

	for index := 0; index < len(values); index++ {
		if yield(values[index]) {
			result = append(result[:index], result[index+1:]...)
			counter++
		}

	}

	return result
}

func IsSorted[T constraints.Ordered](values []T) bool {
	len := len(values)
	previous := values[0]

	for index := 1; index < len; index++ {
		if previous > values[index] {
			return false
		}
		previous = values[index]
	}
	return true
}
