// Good morning! Here's your coding interview problem for today.

// This problem was asked by Microsoft.

// Compute the running median of a sequence of numbers. That is, given a stream of numbers, print out the median of the list so far on each new element.

// Recall that the median of an even-numbered list is the average of the two middle numbers.

// For example, given the sequence [2, 1, 5, 7, 2, 0, 5], your algorithm should print out:

// 2
// 1.5
// 2
// 3.5
// 2
// 2
// 2

package thirtythree

import (
	"container/heap"
	"fmt"
	"math"
)

func ThirtyThree(numbers []int) {

	minHeap := &MinIntHeap{}
	maxHeap := &MaxIntHeap{}
	median := float64(numbers[0])
	heap.Push(minHeap, numbers[0])

	for _, num := range numbers[1:] {
		fmt.Println(median)

		if float64(num) > median {
			heap.Push(minHeap, num)
		} else {
			heap.Push(maxHeap, num)
		}

		//fmt.Println(minHeap)
		//fmt.Println(maxHeap)
		balanceHeaps(minHeap, maxHeap)

		if maxHeap.Len() == minHeap.Len() {
			a := heap.Pop(maxHeap)
			b := heap.Pop(minHeap)

			median = float64(a.(int)+b.(int)) / 2

			heap.Push(maxHeap, a)
			heap.Push(minHeap, b)
		} else if maxHeap.Len() > minHeap.Len() {
			a := heap.Pop(maxHeap)

			median = float64(a.(int))

			heap.Push(maxHeap, a)
		} else {
			a := heap.Pop(minHeap)

			median = float64(a.(int))

			heap.Push(minHeap, a)
		}
	}

	fmt.Println(median)
}

func balanceHeaps(heap_a *MinIntHeap, heap_b *MaxIntHeap) {

	diff := math.Abs(float64(heap_a.Len() - heap_b.Len()))

	if diff == 0 || diff == 1 {
		//This is balanced enough.
	} else if diff == 2 {
		if heap_a.Len() > heap_b.Len() {
			elm := heap.Pop(heap_a)
			heap.Push(heap_b, elm)
		}
	} else {
		panic("How did we get here!!!")
	}

}
