package selectionsort

// SelectionSort takes in array and sorts it by dividing the array into two parts
// It then compares from smallest to largest
func SelectionSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		minIndex := i
		for j := i; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[minIndex], arr[i] = arr[i], arr[minIndex]
	}
	return arr
}
