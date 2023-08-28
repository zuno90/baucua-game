package utils

func FindIndex[T any](slice []T, matchFunc func(T) bool) int {
	for index, element := range slice {
		if matchFunc(element) {
			return index
		}
	}
	return -1 // not found
}

func RemoveElementByIndex[T any](slice []T, index int) []T {
	return append(slice[:index], slice[index+1:]...)
}
