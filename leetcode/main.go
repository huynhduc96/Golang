package main

import (
	"log"
)

/**

 */

func insertIntoBST(root *TreeNode, val int) *TreeNode {

	return root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// test := []int{10, 2, -5}
	// test1 := []int{4, 5, 3, 2, 1}
	// result := asteroidCollision(test)
	// log.Printf("result %+v\n", result)

	root := &TreeNode{Val: 8,
		Left: &TreeNode{Val: 5,
			Left:  &TreeNode{Val: 3},
			Right: &TreeNode{Val: 7},
		},
		Right: &TreeNode{Val: 15,
			Left:  &TreeNode{Val: 14},
			Right: &TreeNode{Val: 18},
		},
	}

	result := insertIntoBST(root, 6)
	log.Printf("result %+v\n", result)

}
