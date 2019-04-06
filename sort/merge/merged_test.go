package merge

import (
	"testing"
)

func TestMerge(t *testing.T) {
	a1 := []int{1, 2, 4, 8, 16}
	a2 := []int{3, 5, 6, 8, 9}

	at := Merge(a1, a2)
	t.Log(at)
}

func TestSort(t *testing.T) {
	a1 := []int{99, 1, 2, 4, 8, 16, 25, 3, 5, 6, 8, 9}

	t.Log("before:", a1)
	at := Sort(a1)
	t.Log("after:", at)
}
