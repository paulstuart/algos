package bubble

// Sort returns a sorted list of ints
func Sort(args ...int) []int {
	// our flag to know if sorting is complete
	swapped := false

	// iterate through the list and if an adjacent
	// pair of values is out order, fix that pair
	for i := 0; i < len(args)-1; i++ {
		if args[i] > args[i+1] {
			args[i], args[i+1] = args[i+1], args[i]
			swapped = true
		}
	}

	// lather, rinse, repeat as needed
	if swapped {
		return Sort(args...)
	}

	return args
}
