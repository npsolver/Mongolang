package global

import (
	"fmt"
)

type TreeNode struct {
	Symbol   *Symbol
	Children []*TreeNode
}

func (tn *TreeNode) AppendChild(child *TreeNode) {
	tn.Children = append([]*TreeNode{child}, tn.Children...)
}

func (tn *TreeNode) printHelper(spaces int) {
	for i := 0; i < spaces; i++ {
		fmt.Print("\t")
	}
	tn.Symbol.Print()
	for _, child := range tn.Children {
		child.printHelper(spaces + 1)
	}
}

func (tn *TreeNode) Print() {
	tn.printHelper(0)
}
