package insertionsort

import (
	"reflect"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	cases := []struct {
		init, exp []int
	}{
		{
			init: []int{1, 2},
			exp:  []int{1, 2},
		},
		{
			init: []int{2, 1},
			exp:  []int{1, 2},
		},
		{
			init: []int{2, 1, 3},
			exp:  []int{1, 2, 3},
		},
		{
			init: []int{2, 1, 3, -1, -2, 10, 7},
			exp:  []int{-2, -1, 1, 2, 3, 7, 10},
		},
	}
	for _, c := range cases {
		got := InsertionSort(c.init)
		if !reflect.DeepEqual(got, c.exp) {
			t.Errorf("init: %v, got array %v, expecting %v", c.init, got, c.exp)
		}
	}
	//arr := []int{2, 3, 1, 4, -1}
	//exp := []int{-1, 1, 2, 3, 4}
	//got := InsertionSort(arr)
	//if !reflect.DeepEqual(got, exp) {
	//}
}
