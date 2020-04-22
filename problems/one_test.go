package problems

import (
	"fmt"
	"testing"
)

func TestOne(t *testing.T) {

	var numberstests = []struct {
		numbers []int
		k       int
		out     bool
	}{
		{[]int{1, 2, 3, 4}, 3, true},
		{[]int{1, 2, 3, 4}, 9, false},
		{[]int{1, 4, 3, 4}, 3, false},
		{[]int{10, 15, 3, 7}, 17, true},
		{[]int{1, 2, 3, 4}, 5, true},
		{[]int{1, 2, 3, 4}, 7, true},
	}

	for _, tt := range numberstests {
		t.Run(fmt.Sprint(tt.numbers), func(t *testing.T) {
			b := One(tt.k, tt.numbers)
			if b != tt.out {
				t.Errorf("got %v, wanted %v", b, tt.out)
			}
		})
	}
}
