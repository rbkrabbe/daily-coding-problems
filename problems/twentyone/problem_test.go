package twentyone

import "testing"

func TestTwentyOne(t *testing.T) {
	max := Twentyone([]Slot{Slot{30, 75}, Slot{0, 50}, Slot{60, 150}})
	if max != 2 {
		t.Errorf("got %v, wanted %v", max, 2)
	}

}
