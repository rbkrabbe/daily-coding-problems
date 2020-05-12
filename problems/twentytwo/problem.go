// Good morning! Here's your coding interview problem for today.

// This problem was asked by Microsoft.

// Given a dictionary of words and a string made up of those words (no spaces), return the original sentence in a list. If there is more than one possible reconstruction, return any of them. If there is no possible reconstruction, then return null.

// For example, given the set of words 'quick', 'brown', 'the', 'fox', and the string "thequickbrownfox", you should return ['the', 'quick', 'brown', 'fox'].

// Given the set of words 'bed', 'bath', 'bedbath', 'and', 'beyond', and the string "bedbathandbeyond", return either ['bed', 'bath', 'and', 'beyond] or ['bedbath', 'and', 'beyond'].

package twentytwo

import (
	"strings"
)

// Twentytwo returns the strings from dict that make up compositeWord, or nil if the word cannot be composed from dict.
func Twentytwo(dict map[string]bool, compositeWord string) []string {

	if len(compositeWord) == 0 {
		return []string{}
	}

	builder := strings.Builder{}
	for _, r := range compositeWord {
		builder.WriteRune(r)

		if _, ok := dict[builder.String()]; ok {
			words := Twentytwo(dict, strings.TrimPrefix(compositeWord, builder.String()))

			if words != nil {
				return append(words, builder.String())
			}
		}
	}

	return nil

}
