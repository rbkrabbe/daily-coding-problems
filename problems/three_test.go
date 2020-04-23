package problems

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThree(t *testing.T) {
	root := Node{"root", &Node{"left", &Node{"left.left", nil, nil}, &Node{"left.right", nil, nil}}, &Node{"right", nil, nil}}

	serialized := ThreeSerialize(&root)
	deserialized := ThreeDeserialize(serialized)

	assert.Equal(t, ThreePrint(&root), ThreePrint(&deserialized))

}
