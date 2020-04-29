// Good morning! Here's your coding interview problem for today.

// This problem was asked by Google.

// A unival tree (which stands for "universal value") is a tree where all nodes under it have the same value.

// Given the root to a binary tree, count the number of unival subtrees.

// For example, the following tree has 5 unival subtrees:

//    0
//   / \
//  1   0
//     / \
//    1   0
//   / \
//  1   1

package problems

// Eight Given the root of a tree, returns the number of unival subtrees.
func Eight(node *Node) int {
	count, _ := unival(node)
	return count
}

// unival Returns the number of unival subtrees below the provided node (including the tree rooted at the node), and a bool indicating whether the node is the root of a unival subtree.
func unival(node *Node) (int, bool) {

	if node == nil {
		return 0, true
	}

	left := node.Left
	right := node.Right

	countLeft, uniLeft := unival(left)
	countRight, uniRight := unival(right)

	totalCount := countLeft + countRight
	uniTotal := false

	if uniLeft && uniRight {
		if left != nil && right != nil {
			if node.Val == left.Val && node.Val == right.Val {
				totalCount++
				uniTotal = true
			}
		} else if left != nil {
			if node.Val == left.Val {
				totalCount++
				uniTotal = true
			}
		} else if right != nil {
			if node.Val == right.Val {
				totalCount++
				uniTotal = true
			}
		} else {
			totalCount++
			uniTotal = true
		}
	}

	return totalCount, uniTotal

}
