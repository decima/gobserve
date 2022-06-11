package gobserve

import "sort"

//sortIntMap converts a int map of any type to a reverse-ordered slice of the same elements.
func sortIntMap[T any](items map[int]T) []T {
	listOfKeys := []int{}

	for k, _ := range items {
		listOfKeys = append(listOfKeys, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(listOfKeys)))

	returnValue := []T{}

	for _, val := range listOfKeys {
		returnValue = append(returnValue, items[val])
	}
	return returnValue
}
