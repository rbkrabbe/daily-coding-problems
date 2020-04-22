// Good morning! Here's your coding interview problem for today.

// This problem was asked by Uber.

// Given an array of integers, return a new array such that each element at index i of the new array is the product of all the numbers in the original array except the one at i.

// For example, if our input was [1, 2, 3, 4, 5], the expected output would be [120, 60, 40, 30, 24]. If our input was [3, 2, 1], the expected output would be [2, 3, 6].

// Follow-up: what if you can't use division?

package problems

// Two Returns a new array such that each element at index i of the new array is the product of all the numbers in the original array except the one at i.
func Two(numbers []int) []int {

	var product int = 1

	for _, n := range numbers {
		product *= n
	}

	result := make([]int, len(numbers))

	for i, n := range numbers {
		result[i] = product / n
	}

	return result
}
