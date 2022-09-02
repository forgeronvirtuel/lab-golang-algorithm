package selectsort

func SelectSort(arr []int, compare func(a, b int) int) []int {
	// loop over the array to sort it using permutation
	for i := 0; i < len(arr); i++ {
		// Search the minimal element of the array
		min := i
		for j := i; j < len(arr); j++ {
			if compare(arr[j], arr[min]) < 0 {
				min = j
			}
		}
		// Permute it with the ith cases
		arr[i], arr[min] = arr[min], arr[i]
	}
	return arr
}
