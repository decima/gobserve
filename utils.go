package gobserve

import "sort"

func sortIntMap[T any](items map[int]T) []T {
	listOfKeys := []int{}

	for k, _ := range items {
		listOfKeys = append(listOfKeys, k)
	}
	sort.Ints(listOfKeys)

	returnValue := []T{}

	for _, val := range listOfKeys {
		returnValue = append(returnValue, items[val])
	}
	return returnValue
}
