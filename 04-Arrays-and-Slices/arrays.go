package slices

// Sum function
func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

// SumAll function
func SumAll(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum)
	result := make([]int, lengthOfNumbers)

	for i, numbers := range numbersToSum {
		result[i] = Sum(numbers)
	}

	return result
}
