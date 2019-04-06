package bubble

import (
	"testing"
)

type tests struct {
	Have, Want []int
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestSort(t *testing.T) {
	tt := []tests{
		{
			Have: []int{9, 3, 24, 2, 72, 55},
			Want: []int{2, 3, 9, 24, 55, 72},
		},
	}
	for _, test := range tt {
		got := Sort(test.Have...)
		if !equal(test.Want, got) {
			t.Fatalf("expected: %v - have: %v\n", test.Want, got)
		}
	}
}
