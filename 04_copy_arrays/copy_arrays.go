package _4_copy_arrays

func ConcatWithAppend(s1, s2 []int) []int {
	var s []int
	for i := 0; i < len(s1); i++ {
		s = append(s, s1[i])
	}
	for i := 0; i < len(s2); i++ {
		s = append(s, s2[i])
	}
	return s
}

func ConcatWithFor(s1, s2 []int) []int {
	size := len(s1) + len(s2)
	s := make([]int, size)

	for i := 0; i < len(s1); i++ {
		s[i] = s1[i]
	}
	for i := 0; i < len(s2); i++ {
		s[i+len(s1)] = s2[i]
	}
	return s
}

func ConcatWithCopy(s1, s2 []int) []int {
	size := len(s1) + len(s2)
	s := make([]int, size)
	copy(s, s1)
	copy(s[len(s1):], s2)
	return s
}
