package piscine

func BTreeIsBinary(root *TreeNode) bool {
	return isBSTHelper(root, "", "")
}

func isBSTHelper(node *TreeNode, min, max string) bool {
	if node == nil {
		return true
	}

	if min != "" && node.Data < min {
		return false
	}

	if max != "" && node.Data >= max {
		return false
	}

	if !isBSTHelper(node.Left, min, node.Data) {
		return false
	}

	if !isBSTHelper(node.Right, node.Data, max) {
		return false
	}

	return true
}
