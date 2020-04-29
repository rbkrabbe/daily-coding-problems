package problems

import (
	"fmt"
	"testing"
)

func TestEight(t *testing.T) {

	var tests = []struct {
		tree   Node
		result int
	}{
		{Node{"root", &Node{"left", &Node{"left.left", nil, nil}, &Node{"left.right", nil, nil}}, &Node{"right", nil, nil}}, 3},
		{Node{"0", &Node{"1", nil, nil}, &Node{"0", &Node{"1", &Node{"1", nil, nil}, &Node{"1", nil, nil}}, &Node{"0", nil, nil}}}, 5},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.tree), func(t *testing.T) {
			count := Eight(&tt.tree)
			if count != tt.result {
				t.Errorf("got %v, wanted %v", count, tt.result)
			}
		})
	}
}
