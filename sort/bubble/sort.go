package bubble

func Sort(args ...int) []int {
	swapped := false
	for i := 0; i < len(args)-1; i++ {
		if args[i] > args[i+1] {
			args[i], args[i+1] = args[i+1], args[i]
			swapped = true
		}
	}
	if swapped {
		return Sort(args...)
	}
	return args
}
