package parser

import (
	"fmt"

	"github.com/npsolver/Mongolang/global"
)

type TreeNode struct {
	symbol   *global.Symbol
	children []*TreeNode
}

func (tn *TreeNode) AppendChild(child *TreeNode) {
	tn.children = append([]*TreeNode{child}, tn.children...)
}

func (tn *TreeNode) printHelper(spaces int) {
	for i := 0; i < spaces; i++ {
		fmt.Print("\t")
	}
	tn.symbol.Print()
	for _, child := range tn.children {
		child.printHelper(spaces + 1)
	}
}

func (tn *TreeNode) Print() {
	tn.printHelper(0)
}
