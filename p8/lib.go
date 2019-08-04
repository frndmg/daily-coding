package p8

import "reflect"

// Node struct
type Node struct {
	Val   interface{}
	Left  *Node
	Right *Node
}

// CountUniversalTree function
func CountUniversalTree(n *Node) int {
	c, _, _ := countUniversalTree(n)
	return c
}

func countUniversalTree(n *Node) (int, interface{}, bool) {
	if n == nil {
		return 0, nil, true
	}

	lc, lv, lu := countUniversalTree(n.Left)
	rc, rv, ru := countUniversalTree(n.Right)

	if lu && ru {
		leftLeaf := lc == 0
		rightLeaf := rc == 0

		childrenEq := reflect.DeepEqual(lv, rv)
		leftEq := reflect.DeepEqual(n.Val, lv)
		rightEq := reflect.DeepEqual(n.Val, rv)

		if leftLeaf && rightLeaf ||
			leftLeaf && rightEq ||
			leftEq && rightLeaf ||
			childrenEq && leftEq {
			return lc + rc + 1, n.Val, true
		}
	}

	return lc + rc, n.Val, false
}
