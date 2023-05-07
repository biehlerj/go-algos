package linear_search

// LinearSearch searches an array of n length looking for the given target
// It does so by going through the array from start to end
// The average performance is O(n)
// It returns the index where the target is or -1 if the target is not found
func LinearSearch(arr []int64, arrLen, target int64) int64 {
	for i := range arr {
		if arr[i] == target {
			return int64(i)
		}
	}
	return -1
}
