package twentytwo

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTwentytwo(t *testing.T) {
	var tests = []struct {
		dict          map[string]bool
		compositeWord string
		result        []string
	}{
		{map[string]bool{"quick": true, "brown": true, "the": true, "fox": true}, "thequickbrownfox", []string{"fox", "brown", "quick", "the"}},
		{map[string]bool{"bed": true, "bath": true, "bedbath": true, "and": true, "beyond": true}, "bedbathandbeyond", []string{"beyond", "and", "bath", "bed"}},
		{map[string]bool{"bed": true, "bath": true, "bey": true, "and": true, "beyond": true}, "bedbathandbeyond", []string{"beyond", "and", "bath", "bed"}},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.compositeWord), func(t *testing.T) {
			words := Twentytwo(tt.dict, tt.compositeWord)
			if !reflect.DeepEqual(words, tt.result) {
				t.Errorf("got %v, wanted %v", words, tt.result)
			}
		})
	}
}
