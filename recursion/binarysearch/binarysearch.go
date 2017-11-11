// Package binarysearch provides binary search algorithm
package binarysearch

import "strings"

// FindStringInSortedSlice searches the slice looking for the specified
// key. The slice has to be sorted according to the lexicographic order
// imposed by strings.Compare(a, b). If the key is found the function
// returns the index in the slice at which that key appears.
// If the key does not exist in the slice, the function returns -1.
func FindStringInSortedSlice(key string, slice []string) int {
	return binarySearch(key, slice, 0, len(slice)-1)
}

// binarySearch does the work for FindStringInSortedSlice. It differs
// taking the upper and lower limit of the search.
func binarySearch(key string, slice []string, low, high int) int {
	if low > high {
		return -1
	}
	mid := (low + high) / 2
	cmp := strings.Compare(key, slice[mid])
	if cmp == 0 {
		return mid
	}
	if cmp < 0 {
		return binarySearch(key, slice, low, mid-1)
	} else {
		return binarySearch(key, slice, mid+1, high)
	}
}
