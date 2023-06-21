package avltree

type AVLTree struct {
	value              int
	rebalanceThreshold int
	left               *AVLTree
	right              *AVLTree
}

func (t *AVLTree) Insert(item int) {
	t.insertNode(item)
	leftHeight, rightHeight := t.getChildrenHeights()
	heightDiff := leftHeight - rightHeight
	if getAbs(heightDiff) >= t.rebalanceThreshold {
		rebalanced := false
		if leftHeight > rightHeight {
			rebalanced = t.left.rebalance()
			if !rebalanced {
				t.rotateRight()
			}
		} else {
			rebalanced = t.right.rebalance()
			if !rebalanced {
				t.rotateLeft()
			}
		}
	}
}

func (t *AVLTree) insertNode(item int) {
	if item >= t.value {
		if t.right == nil {
			t.right = &AVLTree{
				value:              item,
				rebalanceThreshold: t.rebalanceThreshold,
			}
		} else {
			t.right.insertNode(item)
		}
	} else {
		if t.left == nil {
			t.left = &AVLTree{
				value:              item,
				rebalanceThreshold: t.rebalanceThreshold,
			}
		} else {
			t.left.insertNode(item)
		}
	}
}

func (t *AVLTree) getChildrenHeights() (int, int) {
	return t.left.getTreeHeight(), t.right.getTreeHeight()
}

func (t *AVLTree) getTreeHeight() int {
	leftHeight, rightHeight := -1, -1
	if t.left != nil {
		leftHeight = t.left.getTreeHeight()
	}
	if t.right != nil {
		rightHeight = t.right.getTreeHeight()
	}
	return getMax(leftHeight, rightHeight) + 1
}

func (t *AVLTree) rebalance() bool {
	if t.right == nil {
		if t.left.right != nil {
			t.left.rotateLeft()
		}
		t.rotateRight()
		return true
	}
	if t.left == nil {
		if t.right.left != nil {
			t.right.rotateRight()
		}
		t.rotateLeft()
		return true
	}
	return false
}

func (t *AVLTree) rotateLeft() {
	newLeft := *t
	newLeft.right = nil
	if t.right.left != nil {
		newLeft.right = t.right.left
	}
	t.value = t.right.value
	t.right = t.right.right
	t.left = &newLeft
}

func (t *AVLTree) rotateRight() {
	newRight := *t
	newRight.left = nil
	if t.left.right != nil {
		newRight.left = t.left.right
	}
	t.value = t.left.value
	t.left = t.left.left
	t.right = &newRight
}

func getAbs(value int) int {
	if value < 0 {
		return -value
	}
	return value
}

func getMax(left, right int) int {
	if left > right {
		return left
	} else {
		return right
	}
}

func TreesAreEqual(one *AVLTree, two *AVLTree) bool {
	equal := one.value == two.value
	if !equal {
		return false
	}
	if one.left == nil && two.left != nil {
		return false
	}
	if one.left != nil && two.left == nil {
		return false
	}
	if one.right == nil && two.right != nil {
		return false
	}
	if one.right != nil && two.right == nil {
		return false
	}
	if one.left == nil && two.left == nil &&
		one.right == nil && two.right == nil {
		return equal
	}
	if one.left != nil && two.left != nil {
		equal = TreesAreEqual(one.left, two.left)
		if !equal {
			return false
		}
	}
	if one.right != nil && two.right != nil {
		equal = TreesAreEqual(one.right, two.right)
		if !equal {
			return false
		}
	}
	return equal
}
