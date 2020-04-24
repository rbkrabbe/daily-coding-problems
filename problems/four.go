// Good morning! Here's your coding interview problem for today.

// This problem was asked by Stripe.

// Given an array of integers, find the first missing positive integer in linear time and constant space. In other words, find the lowest positive integer that does not exist in the array. The array can contain duplicates and negative numbers as well.

// For example, the input [3, 4, -1, 1] should give 2. The input [1, 2, 0] should give 3.

// You can modify the input array in-place.

package problems

import (
	"fmt"
)

// Four Given an array of integers, find the first missing positive integer in linear time and constant space.
func Four(numbers []int) int {

	numNegative := shiftNegativeLeft(numbers)
	positive := numbers[numNegative:]

	fmt.Println(positive)

	for _, n := range positive {

		if n < 0 {
			n = -n
		}

		if n <= len(positive) && positive[n-1] > 0 {
			positive[n-1] *= -1
		}
	}

	fmt.Println(positive)

	result := 1
	for _, n := range positive {

		if n > 0 {
			break
		}

		result++
	}

	return result
}

func shiftNegativeLeft(numbers []int) int {

	numNegative := 0

	for i := 0; i < len(numbers); i++ {

		if numbers[i] < 1 {
			tmp := numbers[numNegative]
			numbers[numNegative] = numbers[i]
			numbers[i] = tmp

			numNegative++
		}
	}

	return numNegative
}
