package insertionsort

func InsertionSort(arr []int, compare func(a, b int) int) []int {
	for i := 1; i < len(arr); i += 1 {
		v := arr[i]
		j := i - 1
		// shift all sorted elements to the right until we found
		// the place to insert v
		for ; j > -1 && compare(arr[j], v) > 0; j -= 1 {
			arr[j+1] = arr[j]
		}
		arr[j+1] = v
	}
	return arr
}

func SortAppend(arr []int, elt int, compare func(a, b int) int) []int {
	length := len(arr) + 1
	narr := make([]int, length)
	done := false
	for j, i := 0, 0; j < length-1; j, i = j+1, i+1 {
		if compare(elt, arr[i]) <= 0 && !done {
			narr[j] = elt
			j++
			done = true
		}
		narr[j] = arr[i]
	}
	return narr
}
