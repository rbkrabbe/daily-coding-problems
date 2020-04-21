// Good morning! Here's your coding interview problem for today.
// This problem was recently asked by Google.

// Given a list of numbers and a number k, return whether any two numbers from the list add up to k.
// For example, given [10, 15, 3, 7] and k of 17, return true since 10 + 7 is 17.

// Bonus: Can you do this in one pass?

package problems

// One Checks if any two numbers in the numbers array sum to k.
func One(k int, numbers []int) bool {

	memory := map[int]bool{}

	for _, n := range numbers {
		remainder := k - n

		if memory[remainder] {
			return true
		}

		memory[n] = true
	}

	return false
}
