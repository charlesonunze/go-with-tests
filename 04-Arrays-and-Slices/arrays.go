package slices

// Sum function
func Sum(numbers [5]int) int {
	var sum int

	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}

	return sum
}