// Good morning! Here's your coding interview problem for today.

// This problem was asked by Google.

// Given the root to a binary tree, implement serialize(root), which serializes the tree into a string, and deserialize(s), which deserializes the string back into the tree.

// For example, given the following Node class

// class Node:
//     def __init__(self, val, left=None, right=None):
//         self.val = val
//         self.left = left
//         self.right = right

// The following test should pass:

// node = Node('root', Node('left', Node('left.left')), Node('right'))
// assert deserialize(serialize(node)).left.left.val == 'left.left'

package problems

import (
	"encoding/xml"
	"fmt"
)

// Node Represents a node in a tree.
type Node struct {
	Val   string
	Left  *Node
	Right *Node
}

//ThreePrint Returns a string representing the tree rooted at the provided Node.
func ThreePrint(node *Node) string {

	if node == nil {
		return ""
	}

	return fmt.Sprintf("(%s, l%s, r%s)", node.Val, ThreePrint(node.Left), ThreePrint(node.Right))
}

//ThreeSerialize Returns a string representing the tree rooted at the provided Node.
func ThreeSerialize(node *Node) string {

	output, err := xml.MarshalIndent(node, "  ", "    ")

	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	return string(output)
}

//ThreeDeserialize Returns a Node that is the root of the tree represented by the provided string.
func ThreeDeserialize(node string) Node {
	root := Node{}

	if err := xml.Unmarshal([]byte(node), &root); err != nil {
		fmt.Printf("error: %v\n", err)
	}

	return root
}
