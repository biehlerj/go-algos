package quicksort

import "math/rand"

// Quicksort takes an array and sorts it by using the divide and conquer algorithm
// It selects a pivot element and partitions the array into 2 subarrays from that pivot
// It does this by determining whether the element is less than or greater than the pivot
// The subarrays are then sorted recursively
func Quicksort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	left, right := 0, len(arr)-1
	pivot := rand.Int() % len(arr)
	arr[pivot], arr[right] = arr[right], arr[pivot]

	for i := range arr {
		if arr[i] < arr[right] {
			arr[i], arr[left] = arr[left], arr[i]
			left++
		}
	}

	arr[left], arr[right] = arr[right], arr[left]
	Quicksort(arr[:left])
	Quicksort(arr[left+1:])
	return arr
}
