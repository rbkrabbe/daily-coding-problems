package nine

import (
	"fmt"
	"testing"
)

func TestNine(t *testing.T) {

	var tests = []struct {
		numbers []int
		result  int
	}{
		{[]int{2, 4, 6, 2, 5}, 13},
		{[]int{5, 1, 1, 5}, 10},
		{[]int{6, 5, 3}, 9},
		{[]int{6, 5, 3, 4}, 10},
		{[]int{5, 6, 3}, 8},
		{[]int{5, 6, 3, 4}, 10},
		{[]int{2, 4, -6, 2, 5}, 9},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.numbers), func(t *testing.T) {
			count := Problem(tt.numbers)
			if count != tt.result {
				t.Errorf("got %v, wanted %v", count, tt.result)
			}
		})
	}
}
