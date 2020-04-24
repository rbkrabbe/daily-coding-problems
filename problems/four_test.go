package problems

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFour(t *testing.T) {
	var numberstests = []struct {
		numbers []int
		result  int
	}{
		{[]int{3, 4, -1, 1}, 2},
		{[]int{2, 3, -7, 6, 8, 1, -10, 15}, 4},
		{[]int{1, 2, 0}, 3},
	}

	for _, tt := range numberstests {
		t.Run(fmt.Sprint(tt.numbers), func(t *testing.T) {
			b := Four(tt.numbers)
			if !reflect.DeepEqual(b, tt.result) {
				t.Errorf("got %v, wanted %v", b, tt.result)
			}
		})
	}
}
