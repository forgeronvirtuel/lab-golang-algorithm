package insertionsort

import (
	"reflect"
	"testing"
)

func TestInsertionSort(t *testing.T) {
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
		init, exp []int
		cmp       func(a, b int) int
	}{
		{
			init: []int{1, 2},
			exp:  []int{1, 2},
			cmp:  asccmp,
		},
		{
			init: []int{2, 1},
			exp:  []int{1, 2},
			cmp:  asccmp,
		},
		{
			init: []int{2, 1, 3},
			exp:  []int{1, 2, 3},
			cmp:  asccmp,
		},
		{
			init: []int{2, 1, 3, -1, -2, 10, 7},
			exp:  []int{-2, -1, 1, 2, 3, 7, 10},
			cmp:  asccmp,
		},
		{
			init: []int{2, 1, 3, -1, -2, 10, 7},
			exp:  []int{10, 7, 3, 2, 1, -1, -2},
			cmp:  dsccmp,
		},
	}
	for _, c := range cases {
		got := InsertionSort(c.init, c.cmp)
		if !reflect.DeepEqual(got, c.exp) {
			t.Errorf("init: %v, got array %v, expecting %v", c.init, got, c.exp)
		}
	}
}

func TestSortAppend(t *testing.T) {
	asccmp := func(a, b int) int {
		if a > b {
			return 1
		} else if a == b {
			return 0
		}
		return -1
	}
	cases := []struct {
		init_arr []int
		init_elt int
		exp      []int
		cmp      func(a, b int) int
	}{
		{
			init_arr: []int{1, 3},
			init_elt: 2,
			exp:      []int{1, 2, 3},
			cmp:      asccmp,
		},
		{
			init_arr: []int{1, 3, 5, 5, 10, 11, 15},
			init_elt: 16,
			exp:      []int{1, 3, 5, 5, 10, 11, 15, 16},
			cmp:      asccmp,
		},
	}
	for _, c := range cases {
		got := SortAppend(c.init_arr, c.init_elt, c.cmp)
		if !reflect.DeepEqual(got, c.exp) {
			t.Errorf("init_arr: %v, init_elt: %v, got array %v, expecting %v", c.init_elt, c.init_elt, got, c.exp)
		}
	}
}
