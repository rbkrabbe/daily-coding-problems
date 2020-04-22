package problems

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTwo(t *testing.T) {

	var numberstests = []struct {
		numbers []int
		result  []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{120, 60, 40, 30, 24}},
		{[]int{3, 2, 1}, []int{2, 3, 6}},
	}

	for _, tt := range numberstests {
		t.Run(fmt.Sprint(tt.numbers), func(t *testing.T) {
			b := Two(tt.numbers)
			if !reflect.DeepEqual(b, tt.result) {
				t.Errorf("got %v, wanted %v", b, tt.result)
			}
		})
	}
}
