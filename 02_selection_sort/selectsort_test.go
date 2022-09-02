package selectsort

import (
	"reflect"
	"testing"
)

func TestSelectSort(t *testing.T) {
	asccmp := func(a, b int) int {
		if a > b {
			return 1
		} else if a == b {
			return 0
		}
		return -1
	}
	dsccmp := func(a, b int) int {
		if a > b {
			return -1
		} else if a == b {
			return 0
		}
		return 1
	}
	cases := []struct {
		init_arr []int
		exp      []int
		cmp      func(a, b int) int
	}{
		{
			init_arr: []int{1, 3},
			exp:      []int{1, 3},
			cmp:      asccmp,
		},
		{
			init_arr: []int{1, 3, 2},
			exp:      []int{1, 2, 3},
			cmp:      asccmp,
		},
		{
			init_arr: []int{-1, -10, 2, -11, 3, 1, 10, -15},
			exp:      []int{-15, -11, -10, -1, 1, 2, 3, 10},
			cmp:      asccmp,
		},
		{
			init_arr: []int{-1, -10, 2, -11, 3, 1, 10, -15},
			exp:      []int{10, 3, 2, 1, -1, -10, -11, -15},
			cmp:      dsccmp,
		},
	}
	for _, c := range cases {
		got := SelectSort(c.init_arr, c.cmp)
		if !reflect.DeepEqual(got, c.exp) {
			t.Errorf("init_arr: %v, got array %v, expecting %v", c.init_arr, got, c.exp)
		}
	}

}
