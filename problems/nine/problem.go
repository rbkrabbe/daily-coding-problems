// Good morning! Here's your coding interview problem for today.

// This problem was asked by Airbnb.

// Given a list of integers, write a function that returns the largest sum of non-adjacent numbers. Numbers can be 0 or negative.

// For example, [2, 4, 6, 2, 5] should return 13, since we pick 2, 6, and 5. [5, 1, 1, 5] should return 10, since we pick 5 and 5.

// Follow-up: Can you do this in O(N) time and constant space?

package nine

import "fmt"

// Problem Returns the largest sum of non-adjacent numbers from the provided slice.
func Problem(numbers []int) int {

	values := make([]int, len(numbers))
	for i := range numbers {
		values[i] = valueForIndex(numbers, i)
	}

	indices := make([]int, 0)
	for {
		max := 0
		maxIndex := -1

		for i, v := range values {
			if v > max {
				max = v
				maxIndex = i
			}
		}

		if maxIndex == -1 {
			break
		}
		fmt.Println("values: ", values)
		indices = append(indices, maxIndex)
		values[maxIndex] = 0

		if maxIndex > 0 {
			values[maxIndex-1] = 0

			if numbers[maxIndex-1] > 0 && maxIndex > 2 {
				values[maxIndex-3] -= numbers[maxIndex-1]
			}
		}

		if maxIndex < len(values)-1 {
			values[maxIndex+1] = 0

			if numbers[maxIndex+1] > 0 && maxIndex+3 < len(values) {
				values[maxIndex+3] -= numbers[maxIndex+1]
			}
		}

	}
	fmt.Println("indices: ", indices)
	fmt.Println("numbers: ", numbers)
	sum := 0
	for _, i := range indices {
		sum += numbers[i]
	}

	return sum
}

func valueForIndex(numbers []int, index int) int {

	sum := numbers[index]
	if sum < 0 {
		sum = 0
	}

	if index-1 >= 0 && numbers[index-1] > 0 {
		sum -= numbers[index-1]
	}

	if index-2 >= 0 && numbers[index-2] > 0 {
		sum += numbers[index-2]
	}

	if index+1 < len(numbers) && numbers[index+1] > 0 {
		sum -= numbers[index+1]
	}

	if index+2 < len(numbers) && numbers[index+2] > 0 {
		sum += numbers[index+2]
	}

	return sum
}
