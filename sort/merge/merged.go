package merge

import "fmt"

func hey() {
	fmt.Println("vim-go")
}

// Merge merges 2 sorted lists
func Merge(a, b []int) []int {
	var sorted []int
	for len(a) > 0 && len(b) > 0 {
		a1 := a[0]
		b1 := b[0]
		if a1 < b1 {
			sorted = append(sorted, a1)
			a = a[1:]
		} else {
			sorted = append(sorted, b1)
			b = b[1:]
		}
	}
	sorted = append(sorted, a...)
	sorted = append(sorted, b...)

	return sorted
}

// Sort returns a sorted list
func Sort(a []int) []int {
	if len(a) <= 1 {
		return a
	}
	m := len(a) / 2
	left := Sort(a[:m])
	right := Sort(a[m:])
	return Merge(left, right)
}
