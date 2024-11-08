package poker

import "reflect"

func ContainsInArray[T any](needle T, haystack []T) bool {
	for _, v := range haystack {
		if reflect.DeepEqual(v, needle) {
			return true
		}
		// if v == needle {
		// 	return true
		// }
	}
	return false
}
func Combinations[T any](arr []T, size int) [][]T {
	result := [][]T{}
	if size == 0 {
		return append(result, []T{})
	}
	if len(arr) == 0 {
		return result
	}

	first := arr[0]
	rest := arr[1:]

	withFirst := Combinations(rest, size-1)
	for _, comb := range withFirst {
		result = append(result, append([]T{first}, comb...))
	}
	withoutFirst := Combinations(rest, size)
	result = append(result, withoutFirst...)

	return result
}
