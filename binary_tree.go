package main

type node[T any] struct {
	Left  *node[T]
	Right *node[T]
	Value T
}

type BinaryTree[T any] struct {
	Root *node[T]
}

func (biTree *BinaryTree[T]) FindDepth(nd *node[T], count int) int {
	if nd == nil {
		return count
	}

	count++

	var (
		leftHeight  = biTree.FindDepth(nd.Left, count)
		rightHeight = biTree.FindDepth(nd.Right, count)
	)

	if leftHeight > rightHeight {
		return leftHeight
	} else {
		return rightHeight
	}
}
